package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"zocket.example.com/database"
	"zocket.example.com/producer"
)

// Test case: For Http Handler
func TestCreateNewProductHandler(t *testing.T) {
	//instance of db
	DB, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Cannot connect to database", err)
	}
	//instance of rabbitmq
	rabbitMQProducer, err := producer.NewRabbitMQProducer("message_queue")
	if err != nil {
		fmt.Println("Error creating RabbitMQ producer:", err)
		return
	}
	defer rabbitMQProducer.Close()

	myHandlerFuncL := CreateNewProductHandler(rabbitMQProducer, DB)

	server := httptest.NewServer(http.HandlerFunc(myHandlerFuncL))
	reqBody := []byte(`{"name": "Test Product", "price": 90}`)
	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		t.Fatalf("Error making HTTP request: %v", err)
	}
	defer res.Body.Close()
	//checking status code of the request
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}
