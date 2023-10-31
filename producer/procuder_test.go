package producer

import (
	"testing"
)

func TestRabbitMQProducer_PublishMessage(t *testing.T) {

	producer, err := NewRabbitMQProducer("test_queue")
	if err != nil {
		t.Fatalf("Failed to create RabbitMQProducer: %v", err)
	}
	defer producer.Close()

	// Test case: Publish a message successfully
	err = producer.PublishMessage(85)
	if err != nil {
		t.Errorf("Failed to publish a message: %v", err)
	}

}

func TestRabbitMQProducer_Close(t *testing.T) {
	// Create a RabbitMQProducer instance for testing
	producer, err := NewRabbitMQProducer("test_queue")
	if err != nil {
		t.Fatalf("Failed to create RabbitMQProducer: %v", err)
	}

	// Test case: Close the producer's connection and channel successfully
	err = producer.Close()
	if err != nil {
		t.Errorf("Failed to close RabbitMQProducer: %v", err)
	}

}
