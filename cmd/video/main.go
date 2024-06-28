package main

import (
	"bibi/cmd/video/dal"
	"bibi/cmd/video/rpc"
	"bibi/config"
	video "bibi/kitex_gen/video/videohandler"
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
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

var (
	listenAddr string
	EsClient   *elastic.Client
)

func Init() {
	config.Init(constants.VideoServiceName)
	dal.Init()

	InitEs()
	klog.SetLevel(klog.LevelDebug)
	klog.SetLogger(kitexlogrus.NewLogger(kitexlogrus.WithHook(EsHookLog())))
	tracer.InitJaegerTracer(constants.VideoServiceName)
	rpc.InitInteractionRPC()
	rpc.InitUserRPC()
}

func main() {
	Init()

	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
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
	serviceAddr, _ := netpoll.ResolveTCPAddr("tcp", listenAddr)

	svr := video.NewServer(videoHandlerImpl,
		server.WithRegistry(r),
		server.WithServiceAddr(serviceAddr),
		server.WithSuite(opentracing.NewDefaultServerSuite()),
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

func EsHookLog() *eslogrus.ElasticHook {
	hook, err := eslogrus.NewElasticHook(EsClient, config.ElasticSearch.Host, logrus.DebugLevel, constants.ElasticSearchIndexName)
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
