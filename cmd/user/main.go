package main

import (
	"bibi/cmd/user/dal"
	"bibi/cmd/user/rpc"
	"bibi/config"
	user "bibi/kitex_gen/user/userhandler"
	"bibi/pkg/constants"
	"bibi/pkg/tracer"
	"bibi/pkg/utils"
	"bibi/pkg/utils/eslogrus"
	"crypto/tls"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/netpoll"
	elastic "github.com/elastic/go-elasticsearch/v8"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/registry-nacos/registry"
	kopentracing "github.com/kitex-contrib/tracer-opentracing"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	listenAddr string
	lu         = new(LimiterUpdater)
	EsClient   *elastic.Client
)

func Init() {
	config.Init(constants.UserServiceName)
	dal.Init()
	tracer.InitJaegerTracer(constants.UserServiceName)
	InitEs()
	klog.SetLevel(klog.LevelWarn)
	klog.SetLogger(kitexlogrus.NewLogger(kitexlogrus.WithHook(EsHookLog())))
	rpc.Init()
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

	//go pprof.Pprof()
	err = svr.Run()

	if err != nil {
		klog.Error(err.Error())
	}
}

func EsHookLog() *eslogrus.ElasticHook {
	hook, err := eslogrus.NewElasticHook(EsClient, config.ElasticSearch.Host, logrus.WarnLevel, constants.ElasticSearchIndexName)
	if err != nil {
		klog.Warn(err)
	}

	return hook
}

func InitEs() {
	esConn := fmt.Sprintf("http://%s", config.ElasticSearch.Addr)
	cfg := elastic.Config{
		Addresses: []string{esConn},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	client, err := elastic.NewClient(cfg)
	if err != nil {
		klog.Fatal(err)
	}
	EsClient = client
}
