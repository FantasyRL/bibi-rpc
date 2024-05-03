package main

import (
	"bibi/cmd/chat/service"
	"bibi/kitex_gen/chat"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"context"
)

// ChatHandlerImpl implements the last service interface defined in the IDL.
type ChatHandlerImpl struct{}

// MessageRecord implements the ChatHandlerImpl interface.
func (s *ChatHandlerImpl) MessageRecord(ctx context.Context, req *chat.MessageRecordRequest) (resp *chat.MessageRecordResponse, err error) {
	resp = new(chat.MessageRecordResponse)
	switch req.ActionType {
	case 1:
		msgList, count, err := service.NewMessageService(ctx).MessageRecord(req)
		resp.Base = pack.BuildBaseResp(err)
		if err != nil {
			return resp, nil
		}
		resp.MessageCount = count
		resp.Record = service.BuildMessageResp(msgList)
	default:
		resp.Base = pack.BuildBaseResp(errno.ParamError)
	}
	return resp, nil
}

// IsNotReadMessage implements the ChatHandlerImpl interface.
func (s *ChatHandlerImpl) IsNotReadMessage(ctx context.Context, req *chat.IsNotReadMessageRequest) (resp *chat.IsNotReadMessageResponse, err error) {
	resp = new(chat.IsNotReadMessageResponse)
	count, msgList, err := service.NewMessageService(ctx).IsNotReadMessage(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Count = count
	resp.MessageList = msgList
	return resp, nil
}

// MessageSave implements the ChatHandlerImpl interface.
func (s *ChatHandlerImpl) MessageSave(ctx context.Context, req *chat.MessageSaveRequest) (resp *chat.MessageSaveResponse, err error) {
	resp = new(chat.MessageSaveResponse)
	err = service.NewMessageService(ctx).MessageSave(req)
	resp.Base = pack.BuildBaseResp(err)

	return resp, nil
}
