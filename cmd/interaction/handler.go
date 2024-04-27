package main

import (
	"bibi/cmd/interaction/service"
	interaction "bibi/kitex_gen/interaction"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"context"
	"fmt"
)

// InteractionHandlerImpl implements the last service interface defined in the IDL.
type InteractionHandlerImpl struct{}

// LikeAction implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) LikeAction(ctx context.Context, req *interaction.LikeActionRequest) (resp *interaction.LikeActionResponse, err error) {
	resp = new(interaction.LikeActionResponse)
	if (req.VideoId != nil && req.CommentId != nil) || (req.VideoId == nil && req.CommentId == nil) {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	switch req.ActionType {
	case 1:
		err = service.NewInteractionService(ctx).Like(req, req.UserId)
	case 0:
		err = service.NewInteractionService(ctx).DisLike(req, req.UserId)
	default:
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// LikeList implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) LikeList(ctx context.Context, req *interaction.LikeListRequest) (resp *interaction.LikeListResponse, err error) {
	resp = new(interaction.LikeListResponse)
	videoResp, count, err := service.NewInteractionService(ctx).LikeVideoList(req, req.UserId)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.VideoList = videoResp
	resp.VideoCount = &count
	fmt.Println(*videoResp[1])
	return resp, nil
}

// CommentCreate implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) CommentCreate(ctx context.Context, req *interaction.CommentCreateRequest) (resp *interaction.CommentCreateResponse, err error) {
	resp = new(interaction.CommentCreateResponse)
	_, err = service.NewInteractionService(ctx).CommentCreate(req, req.UserId)

	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// CommentDelete implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) CommentDelete(ctx context.Context, req *interaction.CommentDeleteRequest) (resp *interaction.CommentDeleteResponse, err error) {
	resp = new(interaction.CommentDeleteResponse)
	err = service.NewInteractionService(ctx).CommentDelete(req, req.UserId)
	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// GetLikesCountByVideoIdList implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) GetLikesCountByVideoIdList(ctx context.Context, req *interaction.GetLikesCountByVideoIdListRequest) (resp *interaction.GetLikesCountByVideoIdListResponse, err error) {
	resp = new(interaction.GetLikesCountByVideoIdListResponse)
	likeCountList := make([]int64, len(req.VideoId))
	for i, v := range req.VideoId {
		cnt, _ := service.NewInteractionService(ctx).GetVideoLikeById(v)
		likeCountList[i] = cnt
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.LikeCountList = likeCountList
	return resp, nil
}

// GetIsLikeByVideoIdList implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) GetIsLikeByVideoIdList(ctx context.Context, req *interaction.GetIsLikeByVideoIdListRequest) (resp *interaction.GetIsLikeByVideoIdListResponse, err error) {
	resp = new(interaction.GetIsLikeByVideoIdListResponse)
	isLikeResp, err := service.NewInteractionService(ctx).GetIsLikeByVideoIdList(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.IsLikeList = isLikeResp
	return resp, nil
}

// CommentList implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) CommentList(ctx context.Context, req *interaction.CommentListRequest) (resp *interaction.CommentListResponse, err error) {
	resp = new(interaction.CommentListResponse)
	commentResp, count, err := service.NewInteractionService(ctx).CommentList(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.CommentCount = &count
	resp.CommentList = service.BuildCommentsResp(commentResp)
	return resp, nil
}
