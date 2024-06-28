package rpc

import (
	"bibi/kitex_gen/interaction"
	"bibi/kitex_gen/interaction/interactionhandler"
	"bibi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/registry-nacos/resolver"
	kopentracing "github.com/kitex-contrib/tracer-opentracing"
)

var interactionClient interactionhandler.Client

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
		client.WithSuite(kopentracing.NewDefaultClientSuite()),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)

	if err != nil {
		panic(err)
	}

	interactionClient = c
}

func GetLikeCountByIdList(ctx context.Context, req *interaction.GetLikesCountByVideoIdListRequest) (*interaction.GetLikesCountByVideoIdListResponse, error) {
	resp, err := interactionClient.GetLikesCountByVideoIdList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetIsLikeByIdList(ctx context.Context, req *interaction.GetIsLikeByVideoIdListRequest) (*interaction.GetIsLikeByVideoIdListResponse, error) {
	resp, err := interactionClient.GetIsLikeByVideoIdList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
