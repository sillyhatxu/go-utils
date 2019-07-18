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

	MQConfig *Config
}

func (pc ProducerConf) String() string {
	return fmt.Sprintf("{ Exchange : %s, RoutingKey : %s }", pc.Exchange, pc.RoutingKey)
}

func (pc ProducerConf) Send(producer interface{}) error {
	if pc.MQConfig == nil {
		return fmt.Errorf("MQ Config is nil.")
	}
	log.Infof("RabbitMQ ProducerConf{Exchange: %s; RoutingKey: %s; MQConfig: %v;} : %v", pc.Exchange, pc.RoutingKey, pc.MQConfig)
	producerJSON, err := json.Marshal(producer)
	if err != nil {
		return err
	}
	log.Info(string(producerJSON))
	if len(producerJSON) <= 2 {
		//JSON is "{}"
		return errors.New("Struct to json error.")
	}
	conn, err := amqp.Dial(pc.MQConfig.URL)
	if err != nil {
		log.Errorf("Connection [%v] RabbitMQ error.", pc.MQConfig.URL, err)
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
		pc.MQConfig.QueueConfig.Durable,    // durable
		pc.MQConfig.QueueConfig.AutoDelete, // delete when unused
		pc.MQConfig.QueueConfig.Exclusive,  // exclusive
		pc.MQConfig.QueueConfig.NoWait,     // no-wait
		pc.MQConfig.QueueConfig.Arguments,  // arguments
	)
	if err != nil {
		log.Error("Get RabbitMQ queue error.", err)
		return err
	}
	err = ch.Publish(
		pc.Exchange,                       // exchange
		q.Name,                            // routing key
		pc.MQConfig.QueueConfig.Mandatory, // mandatory
		pc.MQConfig.QueueConfig.Immediate, // immediate
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
