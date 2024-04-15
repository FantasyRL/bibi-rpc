// Code generated by Kitex v0.9.1. DO NOT EDIT.

package interactionhandler

import (
	interaction "bibi/kitex_gen/interaction"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	LikeAction(ctx context.Context, req *interaction.LikeActionRequest, callOptions ...callopt.Option) (r *interaction.LikeActionResponse, err error)
	LikeList(ctx context.Context, req *interaction.LikeListRequest, callOptions ...callopt.Option) (r *interaction.LikeListResponse, err error)
	CommentCreate(ctx context.Context, req *interaction.CommentCreateRequest, callOptions ...callopt.Option) (r *interaction.CommentCreateResponse, err error)
	CommentDelete(ctx context.Context, req *interaction.CommentDeleteRequest, callOptions ...callopt.Option) (r *interaction.CommentDeleteResponse, err error)
	GetLikesCountByVideoIdList(ctx context.Context, req *interaction.GetLikesCountByVideoIdListRequest, callOptions ...callopt.Option) (r *interaction.GetLikesCountByVideoIdListResponse, err error)
	GetIsLikeByVideoIdList(ctx context.Context, req *interaction.GetIsLikeByVideoIdListRequest, callOptions ...callopt.Option) (r *interaction.GetIsLikeByVideoIdListResponse, err error)
	CommentList(ctx context.Context, req *interaction.CommentListRequest, callOptions ...callopt.Option) (r *interaction.CommentListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kInteractionHandlerClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kInteractionHandlerClient struct {
	*kClient
}

func (p *kInteractionHandlerClient) LikeAction(ctx context.Context, req *interaction.LikeActionRequest, callOptions ...callopt.Option) (r *interaction.LikeActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LikeAction(ctx, req)
}

func (p *kInteractionHandlerClient) LikeList(ctx context.Context, req *interaction.LikeListRequest, callOptions ...callopt.Option) (r *interaction.LikeListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LikeList(ctx, req)
}

func (p *kInteractionHandlerClient) CommentCreate(ctx context.Context, req *interaction.CommentCreateRequest, callOptions ...callopt.Option) (r *interaction.CommentCreateResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentCreate(ctx, req)
}

func (p *kInteractionHandlerClient) CommentDelete(ctx context.Context, req *interaction.CommentDeleteRequest, callOptions ...callopt.Option) (r *interaction.CommentDeleteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentDelete(ctx, req)
}

func (p *kInteractionHandlerClient) GetLikesCountByVideoIdList(ctx context.Context, req *interaction.GetLikesCountByVideoIdListRequest, callOptions ...callopt.Option) (r *interaction.GetLikesCountByVideoIdListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLikesCountByVideoIdList(ctx, req)
}

func (p *kInteractionHandlerClient) GetIsLikeByVideoIdList(ctx context.Context, req *interaction.GetIsLikeByVideoIdListRequest, callOptions ...callopt.Option) (r *interaction.GetIsLikeByVideoIdListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetIsLikeByVideoIdList(ctx, req)
}

func (p *kInteractionHandlerClient) CommentList(ctx context.Context, req *interaction.CommentListRequest, callOptions ...callopt.Option) (r *interaction.CommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, req)
}