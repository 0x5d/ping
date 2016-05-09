package rabbit

import (
	"fmt"

	"github.com/castillobg/async-ping-pong/brokers"
)

func init() {
	brokers.Register(new(RabbitFactory), "rabbit")
}

type RabbitFactory struct{}

type RabbitAdapter struct {
	Address string
}

func (ra *RabbitAdapter) Start() error {
	return nil
}

func (*RabbitAdapter) Publish(message, topic string) error {
	fmt.Printf("Sent %s to %s.\n", message, topic)
	return nil
}

func (*RabbitFactory) New(address string) (brokers.BrokerAdapter, error) {
	kafkaAdapter := &RabbitAdapter{Address: address}
	return kafkaAdapter, nil
}
