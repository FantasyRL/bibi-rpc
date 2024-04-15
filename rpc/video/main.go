package main

import (
	"bibi/config"
	video "bibi/kitex_gen/video/videohandler"
	"bibi/pkg/constants"
	"bibi/pkg/utils"
	"bibi/rpc/video/dal"
	"bibi/rpc/video/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/netpoll"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var listenAddr string

func Init() {
	config.Init(constants.VideoServiceName)
	dal.Init()

	rpc.InitInteractionRPC()
	rpc.InitUserRPC()
}

func main() {
	Init()

	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		klog.Fatal(err)
	}

	for index, addr := range config.Service.AddrList {
		if ok := utils.AddrCheck(addr); ok {
			listenAddr = addr
			break
		}

		if index == len(config.Service.AddrList)-1 {
			klog.Fatal("not available addr")
		}
	}

	videoHandlerImpl := new(VideoHandlerImpl)
	serviceAddr, err := netpoll.ResolveTCPAddr("tcp", listenAddr)

	svr := video.NewServer(videoHandlerImpl,
		server.WithRegistry(r),
		server.WithServiceAddr(serviceAddr),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: constants.VideoServiceName,
			}),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}))

	err = svr.Run()

	if err != nil {
		klog.Error(err.Error())
	}
}
