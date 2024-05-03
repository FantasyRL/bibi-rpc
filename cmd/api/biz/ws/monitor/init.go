package monitor

import (
	"bibi/cmd/api/biz/rpc"
	"bibi/cmd/api/biz/ws"
	"bibi/kitex_gen/chat"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/websocket"
)

func (manager *ClientManager) Listen() {
	for {
		klog.Info("monitoring")

		select {

		case client := <-Manager.Register:
			klog.Infof("%v:online\n", client.ID)
			Manager.Clients[client.ID] = client //把连接放到用户管理上

			resp := pack.BuildWsBaseResp(errno.WebSocketSuccess)
			_ = client.Socket.WriteMessage(websocket.TextMessage, resp)

		case client := <-Manager.Unregister:
			klog.Infof("%v:offline\n", client.ID)
			resp := pack.BuildWsBaseResp(errno.WebSocketLogoutSuccess)
			_ = client.Socket.WriteMessage(websocket.TextMessage, resp)
			close(client.Send)                 //close chan
			delete(Manager.Clients, client.ID) //delete map

		case broadcast := <-Manager.Broadcast:
			if broadcast.Type == 1 {
				marshalMsg := broadcast.Message
				targetId := broadcast.Client.TargetId

				flag := false
				for id, client := range Manager.Clients {
					if id != targetId {
						continue
					}
					select {
					case client.Send <- marshalMsg: //Write()
						flag = true //成功送信
					default:
						close(client.Send)
						delete(Manager.Clients, client.ID)
					}
				}
				var replyMsg ws.ReplyMsg
				_, _ = replyMsg.UnmarshalMsg(marshalMsg)

				rpcResp, err := rpc.MessageSave(broadcast.Client.Ctx, &chat.MessageSaveRequest{
					TargetId: targetId,
					UserId:   replyMsg.From,
					Content:  replyMsg.Content,
					IsOnline: flag,
				})
				if err != nil {
					klog.Error(err)
				}
				if rpcResp.Base.Code != errno.SuccessCode {
					resp := pack.BuildWsBaseResp(errno.ServiceError)
					_ = broadcast.Client.Socket.WriteMessage(websocket.CloseMessage, resp)
				}
			}
		}
	}
}
