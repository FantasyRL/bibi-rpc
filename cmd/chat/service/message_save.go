package service

import (
	"bibi/cmd/chat/dal/cache"
	"bibi/cmd/chat/dal/db"
	"bibi/cmd/chat/dal/mq"
	"bibi/cmd/chat/service/ws"
	"bibi/kitex_gen/chat"
	"bibi/pkg/errno"
)

func (s *MessageService) MessageSave(req *chat.MessageSaveRequest) error {

	_, err := db.CreateMessage(s.ctx, &db.Message{
		Uid:      req.UserId,
		TargetId: req.TargetId,
		Content:  req.Content,
	})
	if err != nil {
		return err
	}

	if !req.IsOnline {
		//flag==false对方不在线
		//rabbitmq

		marshalMsg, _ := (ws.ReplyMsg{
			Code:    errno.WebSocketSuccessCode,
			From:    req.UserId,
			Content: req.Content,
		}).MarshalMsg(nil)

		if err = mq.NewChatMQ(req.TargetId).Publish(marshalMsg); err != nil {
			return err
		}

		if err = cache.SetMessage(s.ctx, req.TargetId, marshalMsg); err != nil {
			return err
		}
	}
	return nil
}
