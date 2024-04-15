package service

import (
	"bibi/kitex_gen/interaction"
	"bibi/pkg/constants"
	"bibi/pkg/errno"
	"bibi/rpc/interaction/dal/cache"
	"bibi/rpc/interaction/dal/db"
	"errors"
	"gorm.io/gorm"
)

type likeType struct {
	suffix, zset string
	targetId     int64
	dbType       bool //true:video false:comment
}

// 好麻烦，我还是不做了...
//
//todo:isVideoExist;isAuthor(uid:video_id:countSuffix)
func (s *InteractionService) Like(req *interaction.LikeActionRequest, uid int64) error {

	//实现功能复用
	//打包一个结构体吧
	//其实redis.client也要转
	lkType := new(likeType)
	if req.VideoId != nil {
		lkType.suffix = constants.VideoLikeSuffix
		lkType.targetId = *req.VideoId
		lkType.zset = constants.VideoLikeZset
		lkType.dbType = true
	} else {
		lkType.suffix = constants.CommentLikeSuffix
		lkType.targetId = *req.CommentId
		lkType.zset = constants.CommentLikeZset
		lkType.dbType = false
	}
	//用户数据是否存在于redis
	exist, err := cache.IsUserLikeCacheExist(s.ctx, uid, lkType.suffix)
	if err != nil {
		return err
	}
	if !exist {
		if lkType.dbType {
			videoIdList, err := db.GetVideoByUid(s.ctx, uid)
			if err != nil {
				return err
			}
			err = cache.AddLikeList(s.ctx, videoIdList, uid, lkType.suffix)
			if err != nil {
				return err
			}
		}

	}

	//该点赞是否存在
	exist1, err := cache.IsLikeExist(s.ctx, lkType.targetId, uid, lkType.suffix)
	if err != nil {
		return err
	}
	if exist1 {
		return errno.LikeExistError
	}

	//点赞量redis是否过期,若过期则直接存入mysql，未过期则同步视频点赞量
	ok, _, err := cache.GetLikeCount(s.ctx, lkType.zset, lkType.targetId)
	if err != nil {
		return err
	}
	//存在
	if ok {
		//向redis添加用户点赞与视频点赞量
		if err := cache.AddLikeCount(s.ctx, lkType.zset, lkType.targetId, uid, lkType.suffix); err != nil {
			return err
		}

	} else {
		//只添加用户点赞
		if err := cache.AddUserLikeVideo(s.ctx, lkType.targetId, uid, lkType.suffix); err != nil {
			return err
		}
	}

	//检查点赞条目是否存在，存在则更新，不存在则创建
	if lkType.dbType {
		err = db.IsVideoLikeExist(s.ctx, uid, lkType.targetId)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//创建点赞
			return db.VideoLikeCreate(s.ctx, uid, lkType.targetId, 1)
		}
		return db.VideoLikeStatusUpdate(s.ctx, uid, lkType.targetId, 1)
	} else {
		err = db.IsCommentLikeExist(s.ctx, uid, lkType.targetId)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//创建点赞
			return db.CommentLikeCreate(s.ctx, uid, lkType.targetId, 1)
		}
		return db.CommentLikeStatusUpdate(s.ctx, uid, lkType.targetId, 1)
	}

}

func (s *InteractionService) DisLike(req *interaction.LikeActionRequest, uid int64) error {
	lkType := new(likeType)
	if req.VideoId != nil {
		lkType.suffix = constants.VideoLikeSuffix
		lkType.targetId = *req.VideoId
		lkType.zset = constants.VideoLikeZset
		lkType.dbType = true
	} else {
		lkType.suffix = constants.CommentLikeSuffix
		lkType.targetId = *req.CommentId
		lkType.zset = constants.CommentLikeZset
		lkType.dbType = false
	}

	exist, err := cache.IsLikeExist(s.ctx, lkType.targetId, uid, lkType.suffix)
	if err != nil {
		return err
	}
	if !exist {
		if lkType.dbType {
			err = db.IsVideoLikeExist(s.ctx, uid, lkType.targetId)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errno.LikeNotExistError
			}
			if err != nil {
				return err
			}

			if err = db.CheckVideoLikeStatus(s.ctx, uid, lkType.targetId, 0); err == nil {
				return errno.LikeNotExistError
			}
		} else {
			err = db.IsCommentLikeExist(s.ctx, uid, lkType.targetId)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errno.LikeNotExistError
			}
			if err != nil {
				return err
			}

			if err = db.CheckCommentLikeStatus(s.ctx, uid, lkType.targetId, 0); err == nil {
				return errno.LikeNotExistError
			}
		}

	}
	if exist {
		if err = cache.DelVideoLikeCount(s.ctx, lkType.zset, lkType.targetId, uid, lkType.suffix); err != nil {
			return err
		}
	}

	if lkType.dbType {
		return db.VideoLikeStatusUpdate(s.ctx, uid, lkType.targetId, 0)
	} else {
		return db.CommentLikeStatusUpdate(s.ctx, uid, lkType.targetId, 0)
	}
}
