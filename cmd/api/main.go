// Code generated by hertz generator.

package main

import (
	"bibi/cmd/api/biz/mw/jwt"
	"bibi/cmd/api/biz/rpc"
	"bibi/cmd/api/biz/ws/monitor"
	"bibi/config"
	"bibi/pkg/constants"
	"bibi/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/pkg/klog"
)

var listenAddr string

func Init() {
	config.Init(constants.APIServiceName)
	rpc.Init()
	jwt.Init()

}
func main() {
	Init()

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
	h := server.New(
		server.WithHostPorts(listenAddr),
		server.WithStreamBody(true),
		server.WithMaxRequestBodySize(constants.MaxRequestBodySize), //最大字节数
	)

	//websocket
	//NoHijackConnPool 将控制是否使用缓存池来获取/释放劫持连接。
	//如果使用池，将提升内存资源分配的性能，但无法避免二次关闭连接导致的异常。
	h.NoHijackConnPool = true
	go monitor.Manager.Listen()

	register(h)
	h.Spin()
}
