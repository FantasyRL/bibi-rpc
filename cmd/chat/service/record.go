package service

import (
	"bibi/cmd/chat/dal/db"
	"bibi/kitex_gen/chat"
	"bibi/pkg/errno"
	"time"
)

func (s *MessageService) MessageRecord(req *chat.MessageRecordRequest) ([]db.Message, int64, error) {
	ft, _ := time.Parse(time.DateOnly, req.FromTime)
	tt, _ := time.Parse(time.DateOnly, req.ToTime)
	t := tt.Add(time.Hour * 24)
	fts := ft.Unix()
	ts := t.Unix()
	//day:86400
	if fts <= 0 || ts <= 0 || fts >= ts {
		return nil, 0, errno.ParamError
	}
	return db.GetRecordMessagesByTime(s.ctx, req.UserId, req.TargetId, ft, t, int(req.PageNum))
}
