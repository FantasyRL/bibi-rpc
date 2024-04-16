package main

import (
	"bibi/config"
	"bibi/kitex_gen/interaction"
	"bibi/kitex_gen/user"
	video "bibi/kitex_gen/video"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"bibi/rpc/video/dal/db"
	"bibi/rpc/video/rpc"
	"bibi/rpc/video/service"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

// VideoHandlerImpl implements the last service interface defined in the IDL.
type VideoHandlerImpl struct{}

// PutVideo implements the VideoHandlerImpl interface.
func (s *VideoHandlerImpl) PutVideo(ctx context.Context, req *video.PutVideoRequest) (resp *video.PutVideoResponse, err error) {
	resp = new(video.PutVideoResponse)

	videoName, coverName := pack.GenerateName(req.UserId)

	var eg errgroup.Group
	eg.Go(func() error {
		if err != nil {
			return errno.ReadFileError
		}
		err = service.NewVideoService(ctx).UploadCover(req.Cover, coverName)
		if err != nil {
			return errno.UploadFileError
		}
		return nil
	})
	eg.Go(func() error {
		if err != nil {
			return errno.ReadFileError
		}
		err = service.NewVideoService(ctx).UploadVideo(req.VideoFile, videoName)
		if err != nil {
			return errno.UploadFileError
		}
		return nil
	})
	VideoReq := new(db.Video)
	eg.Go(func() error {
		videoUrl := fmt.Sprintf("%s/%s/video/%s", config.OSS.EndPoint, config.OSS.MainDirectory, videoName)
		coverUrl := fmt.Sprintf("%s/%s/video/%s", config.OSS.EndPoint, config.OSS.MainDirectory, coverName)
		VideoReq = &db.Video{
			Uid:      req.UserId,
			Title:    req.Title,
			PlayUrl:  videoUrl,
			CoverUrl: coverUrl,
		}
		_, err = service.NewVideoService(ctx).PutVideo(VideoReq)
		if err != nil {
			return err
		}
		return nil
	})
	if err = eg.Wait(); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	return
}

// ListVideo implements the VideoHandlerImpl interface.
func (s *VideoHandlerImpl) ListVideo(ctx context.Context, req *video.ListUserVideoRequest) (resp *video.ListUserVideoResponse, err error) {
	resp = new(video.ListUserVideoResponse)
	videoResp, count, err := service.NewVideoService(ctx).ListVideo(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	videoIdList := make([]int64, len(*videoResp))
	authorIdList := make([]int64, len(*videoResp))
	//likeCountList := make([]int64, len(*videoResp))
	//authorList := make([]*user.User, len(*videoResp))
	var likeCountList []int64
	var authorList []*user.User
	for i, v := range *videoResp {
		videoIdList[i] = v.ID
		authorIdList[i] = v.Uid
	}
	var eg errgroup.Group
	eg.Go(func() error {
		rpcResp, err := rpc.GetLikeCountByIdList(ctx, &interaction.GetLikesCountByVideoIdListRequest{
			VideoId: videoIdList,
		})
		if err != nil {
			return err
		}
		likeCountList = rpcResp.LikeCountList
		return nil
	})
	eg.Go(func() error {
		rpcResp, err := rpc.UserGetAuthor(ctx, &user.GetAuthorRequest{
			AuthorIdList: authorIdList,
		})
		if err != nil {
			return err
		}
		authorList = rpcResp.AuthorList
		return nil
	})
	//eg.Go(func() error {
	//	rpcResp,err:=rpc.GetIsLikeByIdList(ctx,&interaction.GetIsLikeByVideoIdListRequest{
	//		VideoId: videoIdList,
	//		UserId:
	//	})
	//})
	if err := eg.Wait(); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.VideoList = service.BuildVideoListResp(videoResp, authorList, likeCountList, nil)
	resp.Count = &count
	return resp, nil
}

// SearchVideo implements the VideoHandlerImpl interface.
func (s *VideoHandlerImpl) SearchVideo(ctx context.Context, req *video.SearchVideoRequest) (resp *video.SearchVideoResponse, err error) {
	resp = new(video.SearchVideoResponse)
	videoResp, count, err := service.NewVideoService(ctx).SearchVideo(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}

	videoIdList := make([]int64, len(*videoResp))
	authorIdList := make([]int64, len(*videoResp))
	//likeCountList := make([]int64, len(*videoResp))
	//authorList := make([]*user.User, len(*videoResp))
	var likeCountList []int64
	var authorList []*user.User
	for i, v := range *videoResp {
		videoIdList[i] = v.ID
		authorIdList[i] = v.Uid
	}
	var eg errgroup.Group
	eg.Go(func() error {
		rpcResp, err := rpc.GetLikeCountByIdList(ctx, &interaction.GetLikesCountByVideoIdListRequest{
			VideoId: videoIdList,
		})
		if err != nil {
			return err
		}
		likeCountList = rpcResp.LikeCountList
		return nil
	})
	eg.Go(func() error {
		rpcResp, err := rpc.UserGetAuthor(ctx, &user.GetAuthorRequest{
			AuthorIdList: authorIdList,
		})
		if err != nil {
			return err
		}
		authorList = rpcResp.AuthorList
		return nil
	})
	//eg.Go(func() error {
	//	rpcResp,err:=rpc.GetIsLikeByIdList(ctx,&interaction.GetIsLikeByVideoIdListRequest{
	//		VideoId: videoIdList,
	//		UserId:
	//	})
	//})
	if err := eg.Wait(); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.VideoList = service.BuildVideoListResp(videoResp, authorList, likeCountList, nil)
	resp.Count = &count
	return resp, nil
}

// HotVideo implements the VideoHandlerImpl interface.
func (s *VideoHandlerImpl) HotVideo(ctx context.Context, req *video.HotVideoRequest) (resp *video.HotVideoResponse, err error) {
	// TODO: Your code here...
	return
}
