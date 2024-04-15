package service

import (
	"bibi/kitex_gen/interaction"
	"bibi/pkg/constants"
	"bibi/rpc/interaction/dal/cache"
	"bibi/rpc/interaction/dal/db"
	"errors"
	"gorm.io/gorm"
)

func (s *InteractionService) LikeVideoList(req *interaction.LikeListRequest, uid int64) ([]int64, error) {
	//缓存未过期
	videoIdList, err := cache.GetUserLikeVideos(s.ctx, uid, constants.VideoLikeSuffix)
	if err != nil {
		return nil, err
	}
	if len(videoIdList) != 0 {
		return videoIdList, nil
	}

	//缓存过期
	videoIdList, err = db.GetVideoByUid(s.ctx, uid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	//将mysql数据存入redis缓存
	err = cache.AddLikeList(s.ctx, videoIdList, uid, constants.VideoLikeSuffix)
	if err != nil {
		return nil, err
	}
	return videoIdList, nil
}
