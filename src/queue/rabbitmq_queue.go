package queue

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchange string = "direct"
)

type RabbitMqQueue struct {
	channel *amqp.Channel
	name    string
}

func NewRabbitMqQueue(channel *amqp.Channel, name string) *RabbitMqQueue {
	return &RabbitMqQueue{
		channel: channel,
	}
}

func (this *RabbitMqQueue) Push(message EventMessage) error {
	ctx := context.Background()

	payload, err := json.Marshal(message)

	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		Body: payload,
	}

	return this.channel.PublishWithContext(ctx, exchange, this.name, true, false, msg)
}

func (this *RabbitMqQueue) Pull() (*EventMessage, error) {
	msgChan, err := this.channel.Consume(this.name, "", true, false, false, false, nil)

	if err != nil {
		return nil, err
	}

	msg := <-msgChan

	var message *EventMessage

	err = json.Unmarshal(msg.Body, &message)

	if err != nil {
		return nil, nil
	}

	return message, err
}
