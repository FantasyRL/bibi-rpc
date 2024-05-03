package service

import (
	"bibi/cmd/chat/dal/cache"
	"bibi/kitex_gen/base"
	"bibi/kitex_gen/chat"
)

func (s *MessageService) IsNotReadMessage(req *chat.IsNotReadMessageRequest) (int64, []*base.Message, error) {
	//todo:rabbitmq
	ok, err := cache.IsUserChattedByOthers(s.ctx, req.UserId)
	if err != nil {
		return 0, nil, err
	}
	if ok {
		count, replyMsgs, err := cache.GetMessages(s.ctx, req.UserId)
		if err != nil {
			return 0, nil, err
		}
		msgList := make([]*base.Message, len(replyMsgs))
		for i, msg := range replyMsgs {
			msgList[i] = &base.Message{
				Id:         0,
				TargetId:   msg.Code,
				FromId:     msg.From,
				Content:    msg.Content,
				CreateTime: "",
			}
		}
		return count, msgList, nil
	}
	return 0, nil, nil
}
