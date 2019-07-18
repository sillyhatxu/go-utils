package rabbitmq

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type ConsumerConf struct {
	QueueValue string // @Queue(value='') 相当于  kafka group

	Exchange string

	RoutingKey string // Key 相当于 kafka topic

	MqConfig *Config
}

func (cc ConsumerConf) String() string {
	return fmt.Sprintf("{ QueueValue : %s, Exchange : %s, RoutingKey : %s }", cc.QueueValue, cc.Exchange, cc.RoutingKey)
}

type ConsumerInterface interface {
	MessageDelivery(msg amqp.Delivery)
}

func (cc ConsumerConf) Consumer(ci ConsumerInterface) error {
	log.Infof("RabbitMQ ConsumerConf : %v", cc)
	conn, err := amqp.Dial(cc.MqConfig.URL)
	if err != nil {
		log.Error("Connection RabbitMQ error.", err)
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Error("Get RabbitMQ channel error.", err)
		return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		cc.RoutingKey,                      // name
		cc.MqConfig.QueueConfig.Durable,    // durable
		cc.MqConfig.QueueConfig.AutoDelete, // delete when unused
		cc.MqConfig.QueueConfig.Exclusive,  // exclusive
		cc.MqConfig.QueueConfig.NoWait,     // no-wait
		cc.MqConfig.QueueConfig.Arguments,  // arguments
	)
	if err != nil {
		log.Error("Get RabbitMQ queue error.", err)
		return err
	}
	err = ch.QueueBind(
		q.Name,        // queue name
		cc.RoutingKey, // routing key
		cc.Exchange,   // exchange
		cc.MqConfig.QueueConfig.NoWait,
		cc.MqConfig.QueueConfig.Arguments)
	if err != nil {
		log.Error("RabbitMQ set bind error.", err)
		return err
	}
	msgs, err := ch.Consume(
		q.Name,                            // queue
		"",                                // consumer
		cc.MqConfig.QueueConfig.AutoAck,   // auto-ack
		cc.MqConfig.QueueConfig.Exclusive, // exclusive
		cc.MqConfig.QueueConfig.NoLocal,   // no-local
		cc.MqConfig.QueueConfig.NoWait,    // no-wait
		cc.MqConfig.QueueConfig.Arguments, // args
	)
	if err != nil {
		log.Error("RabbitMQ consume error.", err)
		return nil
	}
	forever := make(chan bool)
	go func() {
		for delivery := range msgs {
			ci.MessageDelivery(delivery)
		}
	}()
	log.Info("Waiting for messages.")
	<-forever
	log.Warningf("RabbitMQ Consumer Exit; ConsumerConf : %v", cc)
	return nil
}
