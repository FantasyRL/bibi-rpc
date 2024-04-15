package main

import (
	interaction "bibi/kitex_gen/interaction"
	"context"
)

// InteractionHandlerImpl implements the last service interface defined in the IDL.
type InteractionHandlerImpl struct{}

// LikeAction implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) LikeAction(ctx context.Context, req *interaction.LikeActionRequest) (resp *interaction.LikeActionResponse, err error) {
	// TODO: Your code here...
	return
}

// LikeList implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) LikeList(ctx context.Context, req *interaction.LikeListRequest) (resp *interaction.LikeListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentCreate implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) CommentCreate(ctx context.Context, req *interaction.CommentCreateRequest) (resp *interaction.CommentCreateResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentDelete implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) CommentDelete(ctx context.Context, req *interaction.CommentDeleteRequest) (resp *interaction.CommentDeleteResponse, err error) {
	// TODO: Your code here...
	return
}

// GetLikesCountByVideoIdList implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) GetLikesCountByVideoIdList(ctx context.Context, req *interaction.GetLikesCountByVideoIdListRequest) (resp *interaction.GetLikesCountByVideoIdListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetIsLikeByVideoIdList implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) GetIsLikeByVideoIdList(ctx context.Context, req *interaction.GetIsLikeByVideoIdListRequest) (resp *interaction.GetIsLikeByVideoIdListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractionHandlerImpl interface.
func (s *InteractionHandlerImpl) CommentList(ctx context.Context, req *interaction.CommentListRequest) (resp *interaction.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
