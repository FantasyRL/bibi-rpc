package db

import (
	"bibi/pkg/constants"
	"context"
	"gorm.io/gorm"
	"time"
)

//go:generate msgp -tests=false -o=chat_msgp.go -io=false
type Message struct {
	ID        int64          `msg:"id"`
	Uid       int64          `msg:"uid"`
	TargetId  int64          `msg:"target"`
	Content   string         `msg:"content"`
	CreatedAt time.Time      `msg:"publish_time"`
	UpdatedAt time.Time      `msg:"-"`
	DeletedAt gorm.DeletedAt `sql:"index" msg:"-"`
}

func CreateMessage(ctx context.Context, message *Message) (*Message, error) {
	if err := DB.WithContext(ctx).Create(message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func GetRecordMessagesByTime(ctx context.Context, uid int64, targetId int64, ft time.Time, tt time.Time, pageNum int) ([]Message, int64, error) {
	msgList := new([]Message)
	var count int64
	if err := DB.WithContext(ctx).Where("uid = ? AND target_id = ? AND created_at <= ? AND created_at >= ?", uid, targetId, tt, ft).
		Count(&count).Limit(constants.PageSize).Offset((pageNum - 1) * constants.PageSize).Find(msgList).Error; err != nil {
		return nil, 0, err
	}
	return *msgList, count, nil
}
