package utils

import (
	"bibi/config"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"net"
	"strconv"
	"strings"
)

func InitMysqlDSN() string {
	return strings.Join([]string{config.Mysql.User, ":", config.Mysql.Password, "@tcp(", config.Mysql.Addr, ")/", config.Mysql.Database, "?charset=utf8mb4&parseTime=True"}, "")

}

func InitRabbitMQDSN() string {
	return strings.Join([]string{"amqp://", config.RabbitMQ.User, ":", config.RabbitMQ.Password, "@", config.RabbitMQ.Addr, "/"}, "")
}

func AddrCheck(addr string) bool {
	l, err := net.Listen("tcp", addr)

	if err != nil {
		return false
	}

	l.Close()

	return true
}

func InitNacos(serviceName string) {
	port, err := strconv.ParseInt(config.Nacos.Port, 10, 64)
	if err != nil {
		klog.Info(err.Error())
	}
	sc := []constant.ServerConfig{{
		IpAddr: config.Nacos.Host,
		Port:   uint64(port),
	}}

	cc := constant.ClientConfig{
		NamespaceId:         "", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		klog.Info(err.Error())
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: serviceName,
		Group:  serviceName,
	})

	if err != nil {
		klog.Info(err.Error())
	}
	fmt.Println(content) //字符串 - yaml
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: serviceName,
		Group:  serviceName,
		OnChange: func(namespace, group, dataId, data string) {
			klog.Info("配置文件发生了变化...")
			klog.Info("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
}
