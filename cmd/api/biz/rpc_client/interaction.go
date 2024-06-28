package rpc_client

import (
	"bibi/kitex_gen/interaction"
	"bibi/kitex_gen/interaction/interactionhandler"
	"context"
	"github.com/kitex-contrib/registry-nacos/resolver"
	opentracing "github.com/kitex-contrib/tracer-opentracing"

	"bibi/pkg/constants"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
)

func InitInteractionRPC() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	c, err := interactionhandler.NewClient(
		constants.InteractionServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(opentracing.NewDefaultClientSuite()),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)

	if err != nil {
		panic(err)
	}

	interactionClient = c
}

func LikeAction(ctx context.Context, req *interaction.LikeActionRequest) (*interaction.LikeActionResponse, error) {
	resp, err := interactionClient.LikeAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func LikedVideoList(ctx context.Context, req *interaction.LikeListRequest) (*interaction.LikeListResponse, error) {
	resp, err := interactionClient.LikeList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CommentCreate(ctx context.Context, req *interaction.CommentCreateRequest) (*interaction.CommentCreateResponse, error) {
	resp, err := interactionClient.CommentCreate(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CommentDelete(ctx context.Context, req *interaction.CommentDeleteRequest) (*interaction.CommentDeleteResponse, error) {
	resp, err := interactionClient.CommentDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CommentList(ctx context.Context, req *interaction.CommentListRequest) (*interaction.CommentListResponse, error) {
	resp, err := interactionClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
