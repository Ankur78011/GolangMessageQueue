package producer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

// Producer interface defines the contract for a message producer.
type Producer interface {
	PublishMessage(message string) error
	Close() error
}

// RabbitMQProducer implements the Producer interface for RabbitMQ.
type RabbitMQProducer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queueName  string
}

// NewRabbitMQProducer creates a new instance of RabbitMQProducer.
func NewRabbitMQProducer(queueName string) (*RabbitMQProducer, error) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	return &RabbitMQProducer{
		connection: connection,
		channel:    channel,
		queueName:  queueName,
	}, nil
}

// PublishMessage publishes a message to the message queue.
func (p *RabbitMQProducer) PublishMessage(prodId int) error {
	productIDString := strconv.Itoa(prodId)
	err := p.channel.Publish(
		"",
		"message_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(productIDString),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	log.Println("Successfully published message to message queue")
	return nil
}

// Close closes the RabbitMQ connection and channel.
func (p *RabbitMQProducer) Close() error {
	if p.channel != nil {
		if err := p.channel.Close(); err != nil {
			return fmt.Errorf("failed to close channel: %v", err)
		}
	}

	if p.connection != nil {
		if err := p.connection.Close(); err != nil {
			return fmt.Errorf("failed to close connection: %v", err)
		}
	}

	return nil
}
