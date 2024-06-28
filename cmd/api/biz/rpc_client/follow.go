package rpc_client

import (
	"bibi/kitex_gen/follow"
	"bibi/kitex_gen/follow/followhandler"
	"bibi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/registry-nacos/resolver"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

func InitFollowRPC() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	c, err := followhandler.NewClient(
		constants.FollowServiceName,
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

	followClient = c
}

func FollowAction(ctx context.Context, req *follow.FollowActionRequest) (*follow.FollowActionResponse, error) {
	resp, err := followClient.FollowAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func FollowerList(ctx context.Context, req *follow.FollowerListRequest) (*follow.FollowerListResponse, error) {
	resp, err := followClient.FollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func FollowingList(ctx context.Context, req *follow.FollowingListRequest) (*follow.FollowingListResponse, error) {
	resp, err := followClient.FollowingList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func FriendList(ctx context.Context, req *follow.FriendListRequest) (*follow.FriendListResponse, error) {
	resp, err := followClient.FriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
