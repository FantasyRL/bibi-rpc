package service

import (
	"bibi/cmd/follow/dal/cache"
	"bibi/cmd/follow/dal/db"
	"bibi/kitex_gen/follow"
	"bibi/pkg/errno"
)

func (s *FollowService) Follow(req *follow.FollowActionRequest) error {
	//redis
	e, err := cache.IsFollowedCacheExist(s.ctx, req.ObjectUid)
	if err != nil {
		return err
	}
	if e {
		e1, err := cache.IsUserFollowExist(s.ctx, req.UserId, req.ObjectUid)
		if err != nil {
			return err
		}
		if e1 {
			return errno.FollowExistError
		}
		err = cache.AddFollower(s.ctx, req.UserId, req.ObjectUid)
		if err != nil {
			return err
		}
	}

	e1, err := db.IsFollowStatus(s.ctx, req.UserId, req.ObjectUid, 1)
	if err != nil {
		return err
	}
	if e1 {
		return errno.FollowExistError
	}

	e2, err := db.IsFollowExist(s.ctx, req.UserId, req.ObjectUid)
	if err != nil {
		return err
	}
	if !e2 {
		return db.CreateFollow(s.ctx, req.UserId, req.ObjectUid, 1)
	}
	return db.UpdateFollowStatus(s.ctx, req.UserId, req.ObjectUid, 1)
}
