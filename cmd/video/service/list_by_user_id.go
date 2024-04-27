package service

import (
	"bibi/cmd/video/dal/db"
	"bibi/cmd/video/rpc"
	"bibi/kitex_gen/base"
	"bibi/kitex_gen/interaction"
	"bibi/kitex_gen/user"
	"bibi/kitex_gen/video"
	"bibi/pkg/errno"
	"golang.org/x/sync/errgroup"
)

func (s *VideoService) ListVideo(req *video.ListUserVideoRequest) (*[]db.Video, int64, []*base.User, []int64, []int64, error) {
	videoList, count, err := db.ListVideosByID(s.ctx, int(req.PageNum), req.UserId)
	if err != nil {
		return nil, 0, nil, nil, nil, err
	}

	videoIdList := make([]int64, len(*videoList))
	authorIdList := make([]int64, len(*videoList))
	var isLikeList []int64
	var likeCountList []int64
	var authorList []*base.User
	for i, v := range *videoList {
		videoIdList[i] = v.ID
		authorIdList[i] = v.Uid
	}
	var eg errgroup.Group
	eg.Go(func() error {
		rpcResp, err := rpc.GetLikeCountByIdList(s.ctx, &interaction.GetLikesCountByVideoIdListRequest{
			VideoId: videoIdList,
		})
		if err != nil {
			return err
		}
		likeCountList = rpcResp.LikeCountList
		return nil
	})
	eg.Go(func() error {
		rpcResp, err := rpc.UserGetUserList(s.ctx, &user.GetUsersRequest{
			UserIdList: authorIdList,
		})
		if err != nil {
			return err
		}
		authorList = rpcResp.UserList
		return nil
	})
	eg.Go(func() error {
		rpcResp, err := rpc.GetIsLikeByIdList(s.ctx, &interaction.GetIsLikeByVideoIdListRequest{
			VideoId: videoIdList,
			UserId:  req.UserId,
		})
		if err != nil {
			return err
		}
		if rpcResp.Base.Code != errno.SuccessCode {
			return errno.NewErrNo(rpcResp.Base.Code, rpcResp.Base.Msg)
		}
		isLikeList = rpcResp.IsLikeList
		return nil
	})
	if err := eg.Wait(); err != nil {
		return nil, 0, nil, nil, nil, err
	}

	return videoList, count, authorList, likeCountList, isLikeList, nil
}
