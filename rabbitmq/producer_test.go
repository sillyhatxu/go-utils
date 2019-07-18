package rabbitmq

import (
	"encoding/json"
	"github.com/sillyhatxu/go-utils/uuid"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

type MqGroupDTO struct {
	Id string `json:"id"`

	Status string `json:"status"`

	OwnerId string `json:"ownerId"`

	OrderId string `json:"orderId"`

	GroupType string `json:"groupType"`

	ProductId string `json:"productId"`

	ProductName string `json:"productName"`
}

func TestJOSN(t *testing.T) {
	producer := "test"
	producerJSON, err := json.Marshal(producer)
	assert.Nil(t, err)
	assert.EqualValues(t, string(producerJSON), producer)
}

func TestProducer(t *testing.T) {
	mqConfig := New("amqp://username:password@127.0.0.1:5672/")
	exchange := "exchange.teste"
	routingKey := "routing.key.test" // Key 相当于 kafka topic
	producer := ProducerConf{Exchange: exchange, RoutingKey: routingKey, MQConfig: mqConfig}
	i := 1
	for {
		err := producer.Send(MqGroupDTO{
			Id:          uuid.UUID(),
			Status:      "Status",
			OwnerId:     "OwnerId",
			OrderId:     "OrderId" + strconv.Itoa(i),
			GroupType:   "GroupType",
			ProductId:   "ProductId",
			ProductName: "ProductName",
		})
		i++
		if err != nil {
			assert.Nil(t, err)
		}
		time.Sleep(5 * time.Second)
	}
}

func TestProducer1(t *testing.T) {
	mqConfig := New("amqp://username:password@127.0.0.1:5672/")
	exchange := "exchange.test"
	routingKey := "test.group.status.change" // Key 相当于 kafka topic
	producer := ProducerConf{Exchange: exchange, RoutingKey: routingKey, MQConfig: mqConfig}
	err := producer.Send(MqGroupDTO{
		Id:          "GB28661706006788428864702515052",
		Status:      "FULL",
		OwnerId:     "U5C8B1E0C2C4E6400019EC5B1",
		OrderId:     "",
		GroupType:   "GROUP_BUY",
		ProductId:   "P7998002600511815841",
		ProductName: "ProductName",
	})
	assert.Nil(t, err)
}
