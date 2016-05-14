package rabbit

import (
	"log"

	"github.com/castillobg/ping/brokers"
	"github.com/streadway/amqp"
)

func init() {
	brokers.Register(new(RabbitFactory), "rabbit")
}

type RabbitFactory struct{}

type RabbitAdapter struct {
	Address    string
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (*RabbitFactory) New(address string) (brokers.BrokerAdapter, func(), error) {
	conn, err := amqp.Dial("amqp://guest:guest@" + address)
	voidFunc := func() {}
	if err != nil {
		return nil, voidFunc, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, voidFunc, err
	}

	cleanup := func() {
		conn.Close()
		ch.Close()
	}
	adapter := &RabbitAdapter{
		Address:    address,
		Channel:    ch,
		Connection: conn,
	}
	return adapter, cleanup, nil
}

func (ra *RabbitAdapter) Listen(queue string, messages chan []byte) error {
	_, err := ra.Channel.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}
	msgs, err := ra.Channel.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	go func() {
		for {
			for d := range msgs {
				log.Printf("Rabbit: Received a message: %s", d.Body)
				messages <- d.Body
			}
		}
	}()
	return nil
}

func (ra *RabbitAdapter) Publish(message, queue string) error {
	err := ra.Channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}
	log.Printf("Rabbit: Sent \"%s\" to queue: \"%s\"", message, queue)
	return nil
}
