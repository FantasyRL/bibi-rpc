package service

import (
	"bibi/kitex_gen/video"
	aliyunoss "bibi/pkg/utils/oss"
	"bibi/rpc/video/dal/db"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

type VideoService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

func NewVideoService(ctx context.Context) *VideoService {
	bucket, err := aliyunoss.OSSBucketCreate()
	if err != nil {
		log.Fatal(err)
	}
	return &VideoService{ctx: ctx, bucket: bucket}
}

func BuildVideoResp(v *db.Video) *video.Video {
	return &video.Video{
		Id:          v.ID,
		Uid:         v.Uid,
		Title:       v.Title,
		PlayUrl:     v.PlayUrl,
		CoverUrl:    v.CoverUrl,
		PublishTime: v.CreatedAt.Format("2006-01-02 15:01:04"),
	}
}

func BuildVideoListResp(videos *[]db.Video, videoLikeCountList []int64, isLikeList []int64) []*video.Video {
	videoListResp := make([]*video.Video, len(*videos))
	for i, v := range *videos {
		videoListResp[i] = BuildVideoResp(&v)
	}
	return videoListResp
}
