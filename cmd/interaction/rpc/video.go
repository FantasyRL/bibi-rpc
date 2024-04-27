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

var videoClient videohandler.Client

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

func VideoGetByIdList(ctx context.Context, req *video.GetVideoByIdListRequest) (*video.GetVideoByIdListResponse, error) {
	resp, err := videoClient.GetVideoByIdList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
