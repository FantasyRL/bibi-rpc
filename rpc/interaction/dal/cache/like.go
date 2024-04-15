package cache

import (
	"bibi/pkg/constants"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
)

// IsUserLikeCacheExist 用户点赞是否存在于redis
func IsUserLikeCacheExist(ctx context.Context, uid int64, suffix string) (bool, error) {
	ok, err := rLike.Exists(ctx, i64ToStr(uid)+suffix).Result()
	if err != nil {
		//错误处理返回啥都一样
		return false, err
	}
	if ok > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// IsLikeExist 用户是否点赞了该video/comment
func IsLikeExist(ctx context.Context, targetId int64, uid int64, suffix string) (bool, error) {
	return rLike.SIsMember(ctx, i64ToStr(uid)+suffix, i64ToStr(targetId)).Result()
}

// AddUserLikeVideo 仅添加用户点赞
func AddUserLikeVideo(ctx context.Context, videoId int64, uid int64, suffix string) error {
	tx := rLike.TxPipeline()
	if err := tx.SAdd(ctx, i64ToStr(uid)+suffix, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, i64ToStr(uid), constants.LikeExpTime).Err(); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// AddLikeCount 添加用户点赞、增加视频点赞量
func AddLikeCount(ctx context.Context, zset string, videoId int64, uid int64, suffix string) error {
	//管线很快，但组装命令过多会导致网络阻塞
	tx := rLike.TxPipeline()
	if err := tx.SAdd(ctx, i64ToStr(uid)+suffix, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	if err := tx.ZIncrBy(ctx, zset, 1, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	//刷新缓存时间
	if err := tx.Expire(ctx, zset, constants.VideoExpTime).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, i64ToStr(uid), constants.LikeExpTime).Err(); err != nil {
		return err
	}

	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// GetLikeCount 获取video/comment点赞量
func GetLikeCount(ctx context.Context, zset string, targetId int64) (bool, int64, error) {
	//获取元素的score
	v, err := rLike.ZScore(ctx, zset, i64ToStr(targetId)).Result()
	if errors.Is(err, redis.Nil) { //已过期
		return false, 0, nil
	}
	if err != nil {
		return true, 114514, err
	}
	cnt := int64(v)
	return true, cnt, nil
}

// DelVideoLikeCount 删除用户点赞、减少视频点赞量
func DelVideoLikeCount(ctx context.Context, zset string, videoId int64, uid int64, suffix string) error {
	tx := rLike.TxPipeline()
	if err := tx.SRem(ctx, i64ToStr(uid)+suffix, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	if err := tx.ZIncrBy(ctx, zset, -1, i64ToStr(videoId)).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, zset, constants.VideoExpTime).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, i64ToStr(uid), constants.LikeExpTime).Err(); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// GetUserLikeVideos 获取用户点赞过的视频ID
func GetUserLikeVideos(ctx context.Context, uid int64, suffix string) ([]int64, error) {
	//SMembers获取所有成员
	vals, err := rLike.SMembers(ctx, i64ToStr(uid)+suffix).Result()
	if err != nil {
		return nil, err
	}

	var videoIdList []int64
	for _, id := range vals {
		vid, _ := strconv.ParseInt(id, 10, 64)
		videoIdList = append(videoIdList, vid)
	}
	return videoIdList, nil
}

// AddLikeList 将用户的所有点赞写入到redis
func AddLikeList(ctx context.Context, targetIdList []int64, uid int64, suffix string) error {
	var err error
	for _, targetId := range targetIdList {
		err = rLike.SAdd(ctx, i64ToStr(uid)+suffix, i64ToStr(targetId)).Err()
	}
	if err != nil {
		return err
	}
	err = rLike.Expire(ctx, i64ToStr(uid), constants.LikeExpTime).Err()
	return err
}

// SetVideoLikeCounts 将视频点赞量写入redis
func SetVideoLikeCounts(ctx context.Context, zset string, videoId int64, likeCount int64) error {
	err := rLike.ZAdd(ctx, zset, redis.Z{
		Score:  float64(likeCount),
		Member: i64ToStr(videoId),
	}).Err()
	if err != nil {
		return err
	}
	err = rLike.Expire(ctx, zset, constants.VideoExpTime).Err()
	return err
}

// ListHotVideo 通过Zset列出点赞最多的视频
func ListHotVideo(ctx context.Context) ([]int64, error) {
	//降序选择前4位点赞最高返回
	res, err := rLike.ZRevRange(ctx, constants.VideoLikeZset, 0, 4).Result()
	if err != nil {
		return nil, err
	}
	var videoIdList []int64
	for _, id := range res {
		vid, _ := strconv.ParseInt(id, 10, 64)
		videoIdList = append(videoIdList, vid)
	}
	return videoIdList, nil
}
