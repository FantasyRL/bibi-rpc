package rpc

import (
	"bibi/config"
	"bibi/kitex_gen/video"
	"bibi/kitex_gen/video/videohandler"
	"bibi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	c, err := videohandler.NewClient(
		constants.VideoServiceName,
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

	videoClient = c
}

func VideoUpload(ctx context.Context, req *video.PutVideoRequest) (*video.PutVideoResponse, error) {
	//rpc client
	resp, err := videoClient.PutVideo(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func UserVideoList(ctx context.Context, req *video.ListUserVideoRequest) (*video.ListUserVideoResponse, error) {
	//rpc client
	resp, err := videoClient.ListVideo(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func VideoUpload(ctx context.Context, req *video.PutVideoReq) (*video.PutVideoResp, error) {
	//rpc client
	resp, err := videoClient.PutVideo(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func VideoUpload(ctx context.Context, req *video.PutVideoReq) (*video.PutVideoResp, error) {
	//rpc client
	resp, err := videoClient.PutVideo(ctx, req)
	//按照逻辑来讲这个err仅用于client出错
	if err != nil {
		return nil, err
	}
	return resp, nil

}
