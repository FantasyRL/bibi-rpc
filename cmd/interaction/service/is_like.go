package service

import (
	"bibi/cmd/interaction/dal/cache"
	"bibi/cmd/interaction/dal/db"
	"bibi/kitex_gen/interaction"
	"bibi/pkg/constants"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (s *InteractionService) GetIsLikeByVideoIdList(req *interaction.GetIsLikeByVideoIdListRequest) ([]int64, error) {
	//缓存未过期
	allVideoIdList, err := cache.GetUserLikeVideos(s.ctx, req.UserId, constants.VideoLikeSuffix)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if allVideoIdList == nil {
		//缓存过期
		allVideoIdList, err = db.GetVideoByUid(s.ctx, req.UserId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

	}

	isLikeResp := make([]int64, len(req.VideoId))
	for i, v := range req.VideoId {
		//也可以用二分，但麻烦点，所以懒了
		isLikeResp[i] = 0
		for j := 0; j < len(allVideoIdList); j++ {
			if allVideoIdList[j] == v {
				isLikeResp[i] = 1
				break
			}
		}
	}
	return isLikeResp, nil
}
