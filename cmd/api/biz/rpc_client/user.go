package rpc_client

import (
	"bibi/kitex_gen/user"
	"bibi/kitex_gen/user/userhandler"
	"bibi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/registry-nacos/resolver"

	//etcd "github.com/kitex-contrib/registry-etcd"
	kopentracing "github.com/kitex-contrib/tracer-opentracing"
)

func GenServiceCBKeyFunc(ri rpcinfo.RPCInfo) string {
	// circuitbreak.RPCInfo2Key returns "$fromServiceName/$toServiceName/$method"
	return circuitbreak.RPCInfo2Key(ri)
}

func InitUserRPC() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}

	//kTracer, kCloser := tracer.InitJaegerTracer("kitex-client")
	//defer kCloser.Close()
	//tracer.InitJaegerTracer("kitex-client")
	//nacosClient, err := nacos.NewClient(nacos.Options{
	//	NamespaceID: constants.APIServiceName,
	//})
	// build a new CBSuite with
	cbs := circuitbreak.NewCBSuite(GenServiceCBKeyFunc)

	c, err := userhandler.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		//client.WithSuite(nacosclient.NewSuite(constants.UserServiceName, constants.APIServiceName, nacosClient)),
		client.WithSuite(kopentracing.NewDefaultClientSuite()),
		//client.WithSuite(kopentracing.NewClientSuite(kTracer, func(c context.Context) string {
		//	endpoint := rpcinfo.GetRPCInfo(c).From()
		//	return endpoint.ServiceName() + "::" + endpoint.Method()
		//})), //jaeger
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
		client.WithCircuitBreaker(cbs), // add CBSuite to the client options
	)
	// update circuit breaker config for a certain key (should be consistent with GenServiceCBKeyFunc)
	// this can be called at any time, and will take effect for following requests
	cbs.UpdateServiceCBConfig("fromServiceName/toServiceName/method", circuitbreak.CBConfig{
		Enable:    true,
		ErrRate:   0.3, // requests will be blocked if error rate >= 30%
		MinSample: 200, // this config takes effect if sampled requests are more than `MinSample`
	})

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

func UserInfo(ctx context.Context, req *user.InfoRequest) (*user.InfoResponse, error) {
	resp, err := userClient.Info(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UserSwitch2FA(ctx context.Context, req *user.Switch2FARequest) (*user.Switch2FAResponse, error) {
	resp, err := userClient.Switch2FA(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UserAvatar(ctx context.Context, req *user.AvatarRequest) (*user.AvatarResponse, error) {
	resp, err := userClient.Avatar(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
