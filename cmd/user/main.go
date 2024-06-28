package main

import (
	"bibi/cmd/user/dal"
	"bibi/config"
	user "bibi/kitex_gen/user/userhandler"
	"bibi/pkg/constants"
	"bibi/pkg/tracer"
	"bibi/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/netpoll"
	"github.com/kitex-contrib/registry-nacos/registry"
	kopentracing "github.com/kitex-contrib/tracer-opentracing"
	"log"
)

var (
	listenAddr string
	lu         = new(LimiterUpdater)
)

//var GloTracer opentracing.Tracer

func Init() {
	config.Init(constants.UserServiceName)
	dal.Init()
	tracer.InitJaegerTracer(constants.UserServiceName)
	//GloTracer = tracer.NewJaegerTracer(constants.UserServiceName, listenAddr)
}

func main() {
	Init()

	//nacos
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
	}

	//获取addr
	for index, addr := range config.Service.AddrList {
		if ok := utils.AddrCheck(addr); ok {
			listenAddr = addr
			break
		}

		if index == len(config.Service.AddrList)-1 {
			klog.Fatal("not available addr")
		}
	}

	userHandlerImpl := new(UserHandlerImpl)
	userCli, err := NewUserClient(listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	serviceAddr, err := netpoll.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	userHandlerImpl.userCli = userCli
	//然而不使用WithServiceAddr方法的话，server还是在监听8888
	//那Impl携带一个Client就没用了

	svr := user.NewServer(userHandlerImpl, // 指定 Registry 与服务基本信息
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: constants.UserServiceName,
			}),
		server.WithSuite(kopentracing.NewDefaultServerSuite()), //jaeger
		//server.WithSuite(kopentracing.NewServerSuite(kTracer, func(c context.Context) string {
		//	endpoint := rpcinfo.GetRPCInfo(c).From()
		//	return endpoint.ServiceName() + "::" + endpoint.Method()
		//})),
		server.WithRegistry(r),
		//server.WithSuite(nacosserver.NewSuite(constants.UserServiceName, nacosClient, cl)),
		server.WithServiceAddr(serviceAddr),
		//server.WithLimit(&limit.Option{
		//	MaxConnections: constants.MaxConnections,
		//	MaxQPS:         constants.MaxQPS,
		//})
		server.WithLimit(
			&limit.Option{
				MaxConnections: constants.MaxConnections,
				MaxQPS:         constants.MaxQPS,
				UpdateControl:  lu.UpdateControl,
			},
		),
	)

	err = svr.Run()

	if err != nil {
		klog.Error(err.Error())
	}
}
