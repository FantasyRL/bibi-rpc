package service

import (
	"bibi/cmd/interaction/dal/cache"
	"bibi/cmd/interaction/dal/db"
	"bibi/kitex_gen/interaction"
	"bibi/kitex_gen/user"
	"bibi/pkg/errno"
	"golang.org/x/sync/errgroup"
)

func (s *InteractionService) CommentCreate(req *interaction.CommentCreateRequest, uid int64) (*db.Comment, error) {
	var eg errgroup.Group
	var err error
	var exist = false
	comment := new(db.Comment)
	commentModel := &interaction.Comment{
		VideoId:  req.VideoId,
		ParentId: req.ParentId,
		Content:  req.Content,
		User: &user.User{
			Id: uid,
		},
	}

	if req.ParentId != nil {
		ok, _ := db.IsParentExist(s.ctx, commentModel)
		if !ok {
			return nil, errno.ParentCommentIsNotExistError
		}
	}

	eg.Go(func() error {
		//若内容完全重复，则删除最早发的那个(其实是懒得再开一个接口了)
		comment, err = db.CreateComment(s.ctx, commentModel)
		if err != nil {
			return err
		}
		return cache.AddVideoComment(s.ctx, comment)
	})

	eg.Go(func() error {
		exist, _, err = cache.GetVideoCommentCount(s.ctx, req.VideoId)
		if err != nil {
			return err
		}
		if exist {
			return cache.IncrVideoCommentCount(s.ctx, req.VideoId)
		}
		return nil
	})

	if err = eg.Wait(); err != nil {
		return nil, err
	}
	if !exist {
		count, err := db.GetCommentCount(s.ctx, req.VideoId)
		if err != nil {
			return nil, err
		}
		err = cache.SetVideoCommentCount(s.ctx, req.VideoId, count)
		if err != nil {
			return nil, err
		}
	}
	return comment, nil
}

func (s *InteractionService) CommentDelete(req *interaction.CommentDeleteRequest, uid int64) error {
	var eg errgroup.Group
	var commentModel = &interaction.Comment{
		Id:      req.CommentId,
		VideoId: req.VideoId,
		User: &user.User{
			Id: uid,
		},
	}
	exist, err := db.IsCommentExist(s.ctx, commentModel)
	if err != nil {
		return err
	}
	if !exist {
		return errno.CommentIsNotExistError
	}
	eg.Go(func() error {

		comment, err := db.DeleteComment(s.ctx, commentModel)
		if err != nil {
			return err
		}
		return cache.DelVideoComment(s.ctx, comment)
	})

	eg.Go(func() error {
		exist, _, err := cache.GetVideoCommentCount(s.ctx, req.VideoId)
		if err != nil {
			return err
		}
		if exist {
			return cache.DecrVideoCommentCount(s.ctx, req.VideoId)
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
