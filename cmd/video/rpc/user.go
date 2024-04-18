package rpc

import (
	"bibi/config"
	"bibi/kitex_gen/user"
	"bibi/kitex_gen/user/userhandler"
	"bibi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userhandler.Client

func InitUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

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

func UserGetAuthor(ctx context.Context, req *user.GetAuthorRequest) (*user.GetAuthorResponse, error) {
	resp, err := userClient.GetAuthor(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
