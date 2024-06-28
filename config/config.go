package config

import (
	"bibi/pkg/utils/remote"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"os"
	"strconv"
)

var (
	OSS           *oss
	Mysql         *mySQL
	Redis         *redis
	Etcd          *etcd
	Server        *server
	Service       *service
	RabbitMQ      *rabbitMQ
	Sender        *email
	ElasticSearch *elasticsearch
	Jaeger        *jaeger
	Nacos         *nacos
	Milvus        *milvus
	runtimeViper  = viper.New()
)

func Init(service string) {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Println("load env error")
	}

	//runtimeViper.SetConfigName("config")
	//runtimeViper.SetConfigType("yaml")
	//runtimeViper.AddConfigPath("./config")
	//if err := runtimeViper.ReadInConfig(); err != nil {
	//	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	//		log.Println("config file not found")
	//	} else {
	//		log.Println("config file was found but another error was produced")
	//	}
	//}
	InitConfigByNacos()
	configMapping(service)

}

func configMapping(serviceName string) {
	c := new(config)
	if err := runtimeViper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}

	Server = &c.Server
	Server.Secret = []byte(runtimeViper.GetString("server.jwt-secret"))

	Mysql = &c.MySQL
	RabbitMQ = &c.RabbitMQ
	Etcd = &c.Etcd
	Redis = &c.Redis
	OSS = &c.OSS
	Sender = &c.Email
	ElasticSearch = &c.ElasticSearch
	Jaeger = &c.Jaeger
	Milvus = &c.Milvus

	addrList := runtimeViper.GetStringSlice("services." + serviceName + ".addr")
	Service = &service{
		Name:     runtimeViper.GetString("services." + serviceName + ".name"),
		AddrList: addrList,
	}

}

func InitConfigByNacos() {
	Nacos = &nacos{
		Host:        os.Getenv("serverAddr"),
		Port:        os.Getenv("serverPort"),
		NamespaceId: os.Getenv("namespaceId"),
		GroupName:   os.Getenv("groupName"),
		DataId:      os.Getenv("dataId"),
		ConfigType:  os.Getenv("type"),
	}
	port, _ := strconv.ParseInt(Nacos.Port, 10, 64)
	remote.SetOptions(&remote.Option{
		Url:         Nacos.Host, // nacos server 多地址需要地址用;号隔开，如 Url: "loc1;loc2;loc3"
		Port:        uint64(port),
		NamespaceId: Nacos.NamespaceId,
		GroupName:   Nacos.GroupName,
		Config: remote.Config{
			DataId: Nacos.DataId,
		},
		Auth: nil, // 如果需要验证登录,需要此参数
	})

	err := runtimeViper.AddRemoteProvider("nacos", Nacos.Host, "")
	if err != nil {
		panic(err)
	}
	runtimeViper.SetConfigType(Nacos.ConfigType)
	_ = runtimeViper.ReadRemoteConfig()
	//_ = runtimeViper.WatchRemoteConfigOnChannel() //异步监听Nacos中的配置变化，如发生配置更改，会直接同步到 viper实例中。

}

func InitTest() {
	Server = &server{
		Version: "debug",
		Name:    "bibi",
		Secret:  []byte("exceed gear"),
	}

	Etcd = &etcd{
		Addr: "127.0.0.1:2379",
	}

	Mysql = &mySQL{
		Addr:     "127.0.0.1:3366",
		User:     "root",
		Password: "114514",
		Database: "bibi_db",
	}

	Redis = &redis{
		Addr: "127.0.0.1:6399",
	}

	RabbitMQ = &rabbitMQ{
		Addr:     "127.0.0.1:5672",
		User:     "guest",
		Password: "guest",
	}

	ElasticSearch = &elasticsearch{
		Addr: "127.0.0.1:9200",
		Host: "127.0.0.1",
	}
}
