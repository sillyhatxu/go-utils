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

	MQConfig *Config
}

func (cc ConsumerConf) String() string {
	return fmt.Sprintf("ConsumerConf{QueueValue: %s, Exchange: %s, RoutingKey: %s ;MQConfig: %v}", cc.QueueValue, cc.Exchange, cc.RoutingKey, cc.MQConfig)
}

type ConsumerInterface interface {
	MessageDelivery(msg amqp.Delivery)
}

func (cc ConsumerConf) Consumer(ci ConsumerInterface) error {
	if cc.MQConfig == nil {
		return fmt.Errorf("MQ Config is nil.")
	}
	log.Infof("RabbitMQ ConsumerConf : %v", cc)
	conn, err := amqp.Dial(cc.MQConfig.URL)
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
		cc.MQConfig.QueueConfig.Durable,    // durable
		cc.MQConfig.QueueConfig.AutoDelete, // delete when unused
		cc.MQConfig.QueueConfig.Exclusive,  // exclusive
		cc.MQConfig.QueueConfig.NoWait,     // no-wait
		cc.MQConfig.QueueConfig.Arguments,  // arguments
	)
	if err != nil {
		log.Error("Get RabbitMQ queue error.", err)
		return err
	}
	err = ch.QueueBind(
		q.Name,        // queue name
		cc.RoutingKey, // routing key
		cc.Exchange,   // exchange
		cc.MQConfig.QueueConfig.NoWait,
		cc.MQConfig.QueueConfig.Arguments)
	if err != nil {
		log.Error("RabbitMQ set bind error.", err)
		return err
	}
	msgs, err := ch.Consume(
		q.Name,                            // queue
		"",                                // consumer
		cc.MQConfig.QueueConfig.AutoAck,   // auto-ack
		cc.MQConfig.QueueConfig.Exclusive, // exclusive
		cc.MQConfig.QueueConfig.NoLocal,   // no-local
		cc.MQConfig.QueueConfig.NoWait,    // no-wait
		cc.MQConfig.QueueConfig.Arguments, // args
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
