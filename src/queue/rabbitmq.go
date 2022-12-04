package queue

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqQueue struct {
	channel *amqp.Channel
	name    string
}

func NewRabbitMqQueue(address string, name string) (*RabbitMqQueue, error) {
	conn, err := amqp.Dial(address)

	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	err = channel.Qos(1, 0, false)

	if err != nil {
		return nil, err
	}

	q := &RabbitMqQueue{
		channel: channel,
	}

	return q, nil
}

func (this *RabbitMqQueue) Push(serializable interface{}) error {
	ctx := context.Background()

	payload, err := json.Marshal(serializable)

	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		Body: payload,
	}

	return this.channel.PublishWithContext(ctx, "direct", this.name, true, false, msg)
}

func (this *RabbitMqQueue) Pull() ([]byte, error) {
	msgChan, err := this.channel.Consume(this.name, "", true, false, false, false, nil)

	if err != nil {
		return nil, err
	}

	msg := <-msgChan
	return msg.Body, err
}
