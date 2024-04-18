package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
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

	runtimeViper = viper.New()
)

func Init(service string) {
	runtimeViper.SetConfigName("config")
	runtimeViper.SetConfigType("yaml")
	runtimeViper.AddConfigPath("./config")
	if err := runtimeViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file not found")
		} else {
			log.Println("config file was found but another error was produced")
		}
	}

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

	addrList := runtimeViper.GetStringSlice("services." + serviceName + ".addr")
	Service = &service{
		Name:     runtimeViper.GetString("services." + serviceName + ".name"),
		AddrList: addrList,
		LB:       runtimeViper.GetBool("services." + serviceName + ".load-balance"), //todo:不知道是啥你也ctrl c？
	}

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
