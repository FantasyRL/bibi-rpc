// Code generated by hertz generator.

package api

import (
	"bibi/cmd/api/biz/rpc"
	"bibi/cmd/api/biz/ws/monitor"
	"bibi/kitex_gen/chat"
	"bibi/pkg/errno"
	"bibi/pkg/pack"

	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/websocket"
	"log"

	api "bibi/cmd/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Chat .
/*// @Param access-token header string false "access-token"// @Param refresh-token header string false "refresh-token"// @router /bibi/message/ws [GET]*/
func Chat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.MessageChatResponse)

	v, _ := c.Get("current_user_id")
	id := v.(int64)
	if id == req.TargetID {
		resp.Base = pack.BuildAPIBaseResp(errno.ParamError)
		c.String(consts.StatusOK, resp.String())
		return
	} //todo:放到register之后

	var upGrader = websocket.HertzUpgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(c *app.RequestContext) bool {
			return true
		},
	}
	err = upGrader.Upgrade(c, func(conn *websocket.Conn) {
		client := &monitor.Client{
			ID:       id,
			TargetId: req.TargetID,
			Socket:   conn,
			Send:     make(chan []byte),
			Ctx:      ctx,
		}

		//将用户注册到用户管理上
		monitor.Manager.Register <- client
		err = client.IfNotReadMessage(id) //rpc
		if err != nil {
			log.Println(err)
			return
		}
		go client.Read()
		go client.Write()
		forever := make(chan bool)
		<-forever //直到conn被关闭才会退出
	})
	if err != nil {
		klog.Info("upgrade:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
}

// MessageRecord .
// @Summary message_record
// @Description get message record
// @Accept json/form
// @Produce json
// @Param target_id query int true "目标id"
// @Param from_time query string true "2024-02-29"
// @Param to_time query string true "2024-03-01"
// @Param action_type query int true "1"
// @Param page_num query int true "1"
// @Param access-token header string false "access-token"
// @Param refresh-token header string false "refresh-token"
// @router /bibi/message/record [GET]
func MessageRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageRecordRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.MessageRecordResponse)

	v, _ := c.Get("current_user_id")
	id := v.(int64)

	rpcResp, err := rpc.MessageRecord(ctx, &chat.MessageRecordRequest{
		TargetId:   req.TargetID,
		FromTime:   req.FromTime,
		ToTime:     req.ToTime,
		ActionType: req.ActionType,
		PageNum:    req.PageNum,
		UserId:     id,
	})
	if err != nil {
		pack.SendRPCFailResp(c, err)
	}
	resp.Base = pack.ConvertToAPIBaseResp(rpcResp.Base)
	if resp.Base.Code != errno.SuccessCode {
		c.JSON(consts.StatusOK, resp)
	}
	resp.MessageCount = rpcResp.MessageCount
	resp.Record = pack.ConvertToAPIMessages(rpcResp.Record)
	c.JSON(consts.StatusOK, resp)
}
