package rpc

import (
	"bibi/config"
	"bibi/kitex_gen/user"
	"bibi/kitex_gen/user/userhandler"
	"bibi/pkg/constants"
	"bibi/pkg/errno"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

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

func UserRegister(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	//rpc client
	resp, err := userClient.Register(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func UserLogin(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	resp, err := userClient.Login(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UserInfo(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
	resp, err := userClient.Info(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.User, nil
}
