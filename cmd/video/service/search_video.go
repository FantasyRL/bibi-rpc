package service

import (
	"bibi/cmd/video/dal/db"
	"bibi/cmd/video/rpc"
	"bibi/kitex_gen/base"
	"bibi/kitex_gen/interaction"
	"bibi/kitex_gen/user"
	"bibi/kitex_gen/video"
	"golang.org/x/sync/errgroup"
)

func (s *VideoService) SearchVideo(req *video.SearchVideoRequest) (*[]db.Video, int64, []*base.User, []int64, []int64, error) {
	videoList, count, err := db.SearchVideo(s.ctx, int(req.PageNum), req.Param)
	if err != nil {
		return nil, 0, nil, nil, nil, err
	}

	videoIdList := make([]int64, len(*videoList))
	authorIdList := make([]int64, len(*videoList))
	//likeCountList := make([]int64, len(*videoResp))
	//authorList := make([]*user.User, len(*videoResp))
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
	//eg.Go(func() error {
	//	rpcResp,err:=rpc.GetIsLikeByIdList(ctx,&interaction.GetIsLikeByVideoIdListRequest{
	//		VideoId: videoIdList,
	//		UserId:,
	//	})
	//})
	if err := eg.Wait(); err != nil {
		return nil, 0, nil, nil, nil, err
	}

	return videoList, count, authorList, likeCountList, nil, nil
}
