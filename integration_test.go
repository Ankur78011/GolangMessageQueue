package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"

	"testing"

	"zocket.example.com/api"

	"zocket.example.com/database"
	"zocket.example.com/producer"
)

func TestIntegration(t *testing.T) {
	DB, err := database.NewDatabase()
	if err != nil {
		log.Println("Cannot connect to database", err)
	}

	rabbitMQProducer, err := producer.NewRabbitMQProducer("message_queue")
	if err != nil {
		log.Println("Error creating RabbitMQ producer:", err)
		return
	}
	defer rabbitMQProducer.Close()
	// instance of hander func
	productHandler := api.CreateNewProductHandler(rabbitMQProducer, DB)
	server := httptest.NewServer(http.HandlerFunc(productHandler))
	reqBody := []byte(`{ 
		"user_id":3,
	   "product_name":"Bat",
	   "product_description":"its a nice bat",
	   "product_images":["https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSGUbHQCMIvIZDFAQIFO9bLVs1hIaIbccxlZA&usqp=CAU"],
	   "product_price":1000
	   }`)
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
