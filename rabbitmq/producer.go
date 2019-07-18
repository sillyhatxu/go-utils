package rabbitmq

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type ProducerConf struct {
	Exchange string

	RoutingKey string // Key 相当于 kafka topic

	MqConfig *Config
}

func (pc ProducerConf) String() string {
	return fmt.Sprintf("{ Exchange : %s, RoutingKey : %s }", pc.Exchange, pc.RoutingKey)
}

func (pc ProducerConf) Send(producer interface{}) error {
	log.Infof("RabbitMQ ProducerConf : %v", pc)
	producerJSON, err := json.Marshal(producer)
	if err != nil {
		return err
	}
	log.Info(string(producerJSON))
	if len(producerJSON) <= 2 {
		//JSON is "{}"
		return errors.New("Struct to json error.")
	}
	conn, err := amqp.Dial(pc.MqConfig.URL)
	if err != nil {
		log.Errorf("Connection [%v] RabbitMQ error.", pc.MqConfig.URL, err)
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
		pc.RoutingKey,                      // name
		pc.MqConfig.QueueConfig.Durable,    // durable
		pc.MqConfig.QueueConfig.AutoDelete, // delete when unused
		pc.MqConfig.QueueConfig.Exclusive,  // exclusive
		pc.MqConfig.QueueConfig.NoWait,     // no-wait
		pc.MqConfig.QueueConfig.Arguments,  // arguments
	)
	if err != nil {
		log.Error("Get RabbitMQ queue error.", err)
		return err
	}
	err = ch.Publish(
		pc.Exchange,                       // exchange
		q.Name,                            // routing key
		pc.MqConfig.QueueConfig.Mandatory, // mandatory
		pc.MqConfig.QueueConfig.Immediate, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(producerJSON),
		})
	if err != nil {
		return err
	}
	log.Info("Send success.")
	return nil
}
