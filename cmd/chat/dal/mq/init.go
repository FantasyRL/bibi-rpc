package mq

import (
	"bibi/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
	"sync"
)

var (
	MQ *RabbitMQ
	Mu sync.Mutex
)

type RabbitMQ struct {
	conn *amqp.Connection
}

type ReplyMsg struct {
	Code    int64  `json:"code" msg:"target"`
	From    int64  `json:"from" msg:"from"`
	Content string `json:"content" msg:"content"`
}

type ChatMQ struct {
	RabbitMQ
	ch        *amqp.Channel
	exchange  string
	queueName string //用于识别未读消息
}

func Init() {

	dial, err := amqp.Dial(utils.InitRabbitMQDSN())
	if err != nil {
		klog.Error(err)
		return
	}
	MQ = &RabbitMQ{
		conn: dial,
	}
}
