package service

import (
	"bibi/cmd/chat/dal/db"
	"bibi/kitex_gen/base"

	"context"
	"time"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService { return &MessageService{ctx: ctx} }

type SendMsg struct {
	Type    int64  `json:"type"`
	Content string `json:"content"`
}

func BuildMessageResp(msgList []db.Message) []*base.Message {
	var msgs []*base.Message
	for _, msg := range msgList {
		msgs = append(msgs, &base.Message{
			Id:         msg.ID,
			TargetId:   msg.TargetId,
			FromId:     msg.Uid,
			Content:    msg.Content,
			CreateTime: msg.CreatedAt.Format(time.RFC3339),
		})
	}
	return msgs
}
