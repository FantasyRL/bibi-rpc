package main

import (
	"bibi/cmd/follow/dal"
	"bibi/cmd/follow/rpc"
	"bibi/config"
	follow "bibi/kitex_gen/follow/followhandler"
	"bibi/pkg/constants"
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
	config.Init(constants.FollowServiceName)
	dal.Init()

	InitEs()
	klog.SetLevel(klog.LevelDebug)
	klog.SetLogger(kitexlogrus.NewLogger(kitexlogrus.WithHook(EsHookLog())))

	rpc.InitUserRPC()
}

func main() {
	Init()
	//注册到etcd
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

	followHandlerImpl := new(FollowHandlerImpl)
	serviceAddr, err := netpoll.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		klog.Fatal(err)
	}

	svr := follow.NewServer(followHandlerImpl, // 指定 Registry 与服务基本信息
		server.WithRegistry(r),
		server.WithServiceAddr(serviceAddr),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: constants.FollowServiceName,
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
