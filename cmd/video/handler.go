package main

import (
	"bibi/cmd/video/dal/db"
	"bibi/cmd/video/service"
	"bibi/config"
	video "bibi/kitex_gen/video"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
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
	videoResp, count, authorList, likeCountList, isLikeList, err := service.NewVideoService(ctx).ListVideo(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}

	resp.VideoList = service.BuildVideoListResp(videoResp, authorList, likeCountList, isLikeList)
	resp.Count = &count
	return resp, nil
}

// SearchVideo implements the VideoHandlerImpl interface.
func (s *VideoHandlerImpl) SearchVideo(ctx context.Context, req *video.SearchVideoRequest) (resp *video.SearchVideoResponse, err error) {
	resp = new(video.SearchVideoResponse)
	videoResp, count, authorList, likeCountList, isLikeList, err := service.NewVideoService(ctx).SearchVideo(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}

	resp.VideoList = service.BuildVideoListResp(videoResp, authorList, likeCountList, isLikeList)
	resp.Count = &count
	return resp, nil
}

// HotVideo implements the VideoHandlerImpl interface.
func (s *VideoHandlerImpl) HotVideo(ctx context.Context, req *video.HotVideoRequest) (resp *video.HotVideoResponse, err error) {
	// TODO: Your code here...
	return
}
