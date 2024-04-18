package service

import (
	"bibi/cmd/interaction/dal/cache"
	"bibi/cmd/interaction/dal/db"
	"bibi/pkg/constants"
	"errors"
	"gorm.io/gorm"
)

func (s *InteractionService) GetVideoLikeById(videoId int64) (int64, error) {
	//redis
	_, likeCount, err := cache.GetLikeCount(s.ctx, constants.VideoLikeZset, videoId)
	if err != nil {
		return 0, err
	}
	if likeCount != 0 {
		return likeCount, nil
	}
	//db
	likeCount, err = db.GetVideoLikeCount(s.ctx, videoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	//存入redis
	if err = cache.SetVideoLikeCounts(s.ctx, constants.VideoLikeZset, videoId, likeCount); err != nil {
		return 0, err
	}
	return likeCount, nil
}
