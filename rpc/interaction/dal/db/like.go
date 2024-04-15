package db

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Like struct {
	ID        int64
	Uid       int64
	VideoId   int64
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}

func CheckVideoLikeStatus(ctx context.Context, uid int64, videoId int64, status int64) error {
	var like Like
	return DBLike.WithContext(ctx).Where("uid = ? AND video_id = ? AND status = ?", uid, videoId, status).
		First(&like).Error
}

func IsVideoLikeExist(ctx context.Context, uid int64, videoId int64) error {
	var like Like
	return DBLike.WithContext(ctx).Where("uid = ? AND video_id = ? ", uid, videoId).
		First(&like).Error
}

func VideoLikeStatusUpdate(ctx context.Context, uid int64, videoId int64, status int64) error {
	return DBLike.WithContext(ctx).Where("uid = ? AND video_id = ? ", uid, videoId).
		Update("status", status).Error
}

func VideoLikeCreate(ctx context.Context, uid int64, videoId int64, status int64) error {
	var like = &Like{
		Uid:     uid,
		VideoId: videoId,
		Status:  status,
	}

	return DBLike.WithContext(ctx).Create(like).Error
}

func GetVideoByUid(ctx context.Context, uid int64) ([]int64, error) {
	likes := new([]Like)
	if err := DBLike.WithContext(ctx).Where("uid = ? AND status = ?", uid, 1).Find(likes).Error; err != nil {
		return nil, err
	}
	var videoIdList []int64
	for _, id := range *likes {
		videoIdList = append(videoIdList, id.VideoId)
	}
	return videoIdList, nil
}

func GetVideoLikeCount(ctx context.Context, videoId int64) (count int64, err error) {
	if err = DBLike.WithContext(ctx).Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
