package service

import (
	"bibi/cmd/video/dal/db"
	"bibi/config"
	"bibi/kitex_gen/base"
	aliyunoss "bibi/pkg/utils/oss"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

type VideoService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

func NewVideoService(ctx context.Context) *VideoService {
	if config.OSS == nil {
		return &VideoService{ctx: ctx, bucket: nil}
	}
	bucket, err := aliyunoss.OSSBucketCreate()
	if err != nil {
		log.Fatal(err)
	}
	return &VideoService{ctx: ctx, bucket: bucket}
}

func BuildVideoResp(v *db.Video, author *base.User, likeCount int64) *base.Video {
	return &base.Video{
		Id:          v.ID,
		Uid:         v.Uid,
		Author:      author,
		Title:       v.Title,
		PlayUrl:     v.PlayUrl,
		CoverUrl:    v.CoverUrl,
		LikeCount:   likeCount,
		PublishTime: v.CreatedAt.Format("2006-01-02 15:01:04"),
	}
}

func BuildVideoListResp(videos *[]db.Video, authorList []*base.User, videoLikeCountList []int64, isLikeList []int64) []*base.Video {
	videoListResp := make([]*base.Video, len(*videos))
	for i, v := range *videos {
		videoListResp[i] = BuildVideoResp(&v, authorList[i], videoLikeCountList[i])
	}
	return videoListResp
}
