package rpc

import (
	"bibi/config"
	"bibi/kitex_gen/chat"
	"bibi/kitex_gen/chat/chathandler"
	"bibi/pkg/constants"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitChatRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	c, err := chathandler.NewClient(
		constants.ChatServiceName,
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

	chatClient = c
}

func MessageSave(ctx context.Context, req *chat.MessageSaveRequest) (*chat.MessageSaveResponse, error) {
	resp, err := chatClient.MessageSave(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func IsNotReadMessage(ctx context.Context, req *chat.IsNotReadMessageRequest) (*chat.IsNotReadMessageResponse, error) {
	resp, err := chatClient.IsNotReadMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func MessageRecord(ctx context.Context, req *chat.MessageRecordRequest) (*chat.MessageRecordResponse, error) {
	resp, err := chatClient.MessageRecord(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
