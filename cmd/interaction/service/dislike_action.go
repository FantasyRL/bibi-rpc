package service

import (
	"bibi/cmd/interaction/dal/cache"
	"bibi/cmd/interaction/dal/db"
	"bibi/kitex_gen/interaction"
	"bibi/pkg/constants"
	"bibi/pkg/errno"
	"errors"
	"gorm.io/gorm"
)

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
