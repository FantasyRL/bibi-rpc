package rpc

import (
	"bibi/kitex_gen/user"
	"bibi/kitex_gen/user/userhandler"
	"bibi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

var userClient userhandler.Client

func InitUserRPC() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	c, err := userhandler.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)

	if err != nil {
		panic(err)
	}

	userClient = c
}

func UserGetUserList(ctx context.Context, req *user.GetUsersRequest) (*user.GetUsersResponse, error) {
	resp, err := userClient.GetUserList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
