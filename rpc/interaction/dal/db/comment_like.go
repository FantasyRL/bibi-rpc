package db

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type CommentLike struct {
	Id        int64
	Uid       int64
	CommentId int64
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CheckCommentLikeStatus(ctx context.Context, uid int64, commentId int64, status int64) error {
	var like CommentLike
	return DBCommentLike.WithContext(ctx).Where("uid = ? AND comment_id = ? AND status = ?", uid, commentId, status).
		First(&like).Error
}

func IsCommentLikeExist(ctx context.Context, uid int64, commentId int64) error {
	var like CommentLike
	return DBCommentLike.WithContext(ctx).Where("uid = ? AND comment_id = ? ", uid, commentId).
		First(&like).Error
}

func CommentLikeStatusUpdate(ctx context.Context, uid int64, commentId int64, status int64) error {
	return DBCommentLike.WithContext(ctx).Where("uid = ? AND comment_id = ? ", uid, commentId).
		Update("status", status).Error
}

func CommentLikeCreate(ctx context.Context, uid int64, commentId int64, status int64) error {
	var like = &CommentLike{
		Uid:       uid,
		CommentId: commentId,
		Status:    status,
	}

	return DBCommentLike.WithContext(ctx).Create(like).Error
}

func GetCommentLikeCount(ctx context.Context, commentId int64) (count int64, err error) {
	if err = DBCommentLike.WithContext(ctx).Where("comment_id = ?", commentId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
