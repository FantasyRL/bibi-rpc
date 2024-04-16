package service

import (
	"bibi/kitex_gen/user"
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

func BuildVideoResp(v *db.Video, author *user.User, likeCount int64) *video.Video {
	return &video.Video{
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

func BuildVideoListResp(videos *[]db.Video, authorList []*user.User, videoLikeCountList []int64, isLikeList []int64) []*video.Video {
	videoListResp := make([]*video.Video, len(*videos))
	for i, v := range *videos {
		videoListResp[i] = BuildVideoResp(&v, authorList[i], videoLikeCountList[i])
	}
	return videoListResp
}
