package mq

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
	"strconv"
)

func NewChatMQ(uid int64) (cli *ChatMQ) {
	cli = &ChatMQ{
		RabbitMQ:  *MQ,
		queueName: strconv.FormatInt(uid, 10),
	}

	ch, err := cli.conn.Channel()
	if err != nil {
		klog.Error(err)
	}
	cli.ch = ch
	return
}

func (c *ChatMQ) Publish(marshalMsg []byte) error {
	q, err := c.ch.QueueDeclare(
		c.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = c.ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        marshalMsg,
		})
	if err != nil {
		return err
	}

	return nil
}

func (c *ChatMQ) Consumer() {
	defer c.conn.Close()
	defer c.ch.Close()

	q, err := c.ch.QueueDeclare(
		c.queueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		klog.Error(err)
		return
	}

	msgs, err := c.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		klog.Error(err)
		return
	}

	//forever := make(chan bool)
	c.Receive(msgs)
	//<-forever
}

func (c *ChatMQ) Receive(msgs <-chan amqp.Delivery) {
	for req := range msgs {
		replyMsg := new(ReplyMsg)
		_, err := replyMsg.UnmarshalMsg(req.Body)
		if err != nil {
			klog.Error(err)
			continue
		}
		fmt.Println(replyMsg)
	}
}
