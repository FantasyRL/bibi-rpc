package service

import (
	"bibi/kitex_gen/interaction"
	"bibi/rpc/interaction/dal/cache"
	"bibi/rpc/interaction/dal/db"
	"errors"
	"gorm.io/gorm"
)

func (s *InteractionService) CommentList(req *interaction.CommentListRequest) ([]db.Comment, int64, error) {
	commentCache, err := cache.GetVideoComments(s.ctx, req.VideoId)
	if err != nil {
		return nil, 0, err
	}
	exist, countCache, err := cache.GetVideoCommentCount(s.ctx, req.VideoId)
	if err != nil {
		return nil, 0, err
	}
	if exist && len(commentCache) != 0 {
		return commentCache, countCache, nil
	}
	if !exist && len(commentCache) != 0 {
		count, err := db.GetCommentCount(s.ctx, req.VideoId)
		if err != nil {
			return nil, 0, err
		}
		return commentCache, count, nil
	}
	comments, count, err := db.GetCommentsByVideoID(s.ctx, req.VideoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}
	if err != nil {
		return nil, 0, err
	}
	//设置缓存
	if err := cache.SetVideoComments(s.ctx, comments, req.VideoId); err != nil {
		return nil, 0, err
	}
	return comments, count, nil
}
