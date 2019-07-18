package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type QueueConfig struct {
	Arguments  amqp.Table
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Mandatory  bool
	Immediate  bool
	AutoAck    bool
	NoLocal    bool
}

type Config struct {
	URL         string
	QueueConfig QueueConfig
}

func (c Config) String() string {
	return fmt.Sprintf("Config{URL: %v; QueueConfig: %+v}", c.URL, c.QueueConfig)
}

func New(url string) *Config {
	return &Config{
		URL: url,
		QueueConfig: QueueConfig{
			Arguments: amqp.Table{
				"name":  "x-message-ttl",
				"value": 3600000,
				"type":  "java.lang.Long",
			},
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Mandatory:  false,
			Immediate:  false,
			AutoAck:    true,
			NoLocal:    false,
		},
	}
}
