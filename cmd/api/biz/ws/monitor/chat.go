package monitor

import (
	"bibi/cmd/api/biz/rpc_client"
	"bibi/cmd/api/biz/ws"
	"bibi/kitex_gen/chat"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/websocket"
)

func (c *Client) Read() {
	defer func() { //闭包
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		sendMsg := new(ws.SendMsg)
		err := c.Socket.ReadJSON(sendMsg) // 接收消息
		if err != nil && !errors.As(err, &websocket.ErrCloseSent) {
			klog.Info(errno.ParamErrMsg + ":" + err.Error())
			break
		}

		if len(sendMsg.Content) == 0 || len(sendMsg.Content) > 2000 {
			resp := pack.BuildWsBaseResp(errno.CharacterBeyondLimitError)
			_ = c.Socket.WriteMessage(websocket.TextMessage, resp)
			break
		}

		if sendMsg.Type == 1 {
			marshalMsg, _ := (ws.ReplyMsg{
				Code:    errno.WebSocketSuccessCode,
				From:    c.ID,
				Content: sendMsg.Content,
			}).MarshalMsg(nil)
			Manager.Broadcast <- &Broadcast{ //传到broadcast来发给target user
				Client:  c,
				Message: marshalMsg,
				Type:    sendMsg.Type,
			}
		}
	}

}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case marshalMsg, ok := <-c.Send:
			if !ok {
				resp := pack.BuildWsBaseResp(errno.WebSocketError)
				_ = c.Socket.WriteMessage(websocket.CloseMessage, resp)
				return
			}
			var replyMsg ws.ReplyMsg
			_, _ = replyMsg.UnmarshalMsg(marshalMsg)
			resp, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, resp)

		}
	}
}

func (c *Client) IfNotReadMessage(uid int64) error {
	rpcResp, err := rpc_client.IsNotReadMessage(c.Ctx, &chat.IsNotReadMessageRequest{
		UserId: uid,
	})
	if err != nil {
		return err
	}
	if rpcResp.Count != 0 {
		countMsg := fmt.Sprintf("你有%v条未读消息", rpcResp.Count)
		_ = c.Socket.WriteMessage(websocket.TextMessage, []byte(countMsg))
		for _, replyMsg := range rpcResp.MessageList {
			resp, _ := sonic.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, resp)
		}
	}
	return nil
}
