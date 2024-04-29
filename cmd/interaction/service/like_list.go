package service

import (
	"bibi/cmd/interaction/dal/cache"
	"bibi/cmd/interaction/dal/db"
	"bibi/cmd/interaction/rpc"
	"bibi/kitex_gen/base"
	"bibi/kitex_gen/interaction"
	"bibi/kitex_gen/video"
	"bibi/pkg/constants"
	"bibi/pkg/errno"
	"errors"
	"gorm.io/gorm"
)

func (s *InteractionService) LikeVideoList(req *interaction.LikeListRequest, uid int64) ([]*base.Video, int64, error) {
	//缓存未过期
	allVideoIdList, err := cache.GetUserLikeVideos(s.ctx, uid, constants.VideoLikeSuffix)
	if err != nil {
		return nil, 0, err
	}

	//缓存过期
	if allVideoIdList == nil {
		allVideoIdList, err = db.GetVideoByUid(s.ctx, uid)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		if err != nil {
			return nil, 0, err
		}
		//将mysql数据存入redis缓存
		err = cache.AddLikeList(s.ctx, allVideoIdList, uid, constants.VideoLikeSuffix)
		if err != nil {
			return nil, 0, err
		}
	}

	length := len(allVideoIdList)
	var videoIdList []int64
	if length <= int(req.PageNum-1)*constants.PageSize || int(req.PageNum-1)*constants.PageSize < 0 {
		return nil, 0, nil
	} else {
		fst := int(req.PageNum-1) * constants.PageSize
		for i := fst; i < fst+constants.PageSize && i < length; i++ {
			videoIdList = append(videoIdList, allVideoIdList[i])
		}
	}

	rpcResp, err := rpc.VideoGetByIdList(s.ctx, &video.GetVideoByIdListRequest{
		VideoIdList: videoIdList,
		UserId:      req.UserId,
	})
	if err != nil {
		return nil, 0, err
	}
	if rpcResp.Base.Code != errno.SuccessCode {
		return nil, 0, errno.NewErrNo(rpcResp.Base.Code, rpcResp.Base.Msg)
	}
	videoCount := rpcResp.Count
	videosResp := make([]*base.Video, 0)
	//videosResp := make([]*base.Video, len(rpcResp.VideoList))
	//for i, v := range rpcResp.VideoList {
	//	videosResp[i] = v
	//}
	copy(videosResp, rpcResp.VideoList)
	return videosResp, videoCount, nil
}
