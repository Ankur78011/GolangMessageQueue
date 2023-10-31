package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"zocket.example.com/api"
	"zocket.example.com/consumer"
	"zocket.example.com/database"
	"zocket.example.com/producer"
)

func TestIntegration(t *testing.T) {
	DB, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Cannot connect to database", err)
	}

	rabbitMQProducer, err := producer.NewRabbitMQProducer("message_queue")
	if err != nil {
		fmt.Println("Error creating RabbitMQ producer:", err)
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

	//We have prodcut id = 88 of bat in database

	//we will recevice 88 as a prodcut id from consumer
	err = consumer.ConsumerFunction(strconv.Itoa(88), DB)
	if err != nil {
		t.Error("Consumerfunction not working")
	}
	expected := []string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSGUbHQCMIvIZDFAQIFO9bLVs1hIaIbccxlZA&usqp=CAU"}
	expectedString := strings.Join(expected, ", ")

	arrayOfImageUrl, err := DB.GetProductFromProductId(88)
	if err != nil {
		t.Error("Failed to get urls of images")
	}
	arrayOfImageUrlString := strings.Join(arrayOfImageUrl, ",")

	if arrayOfImageUrlString != expectedString {
		t.Error("Failed to get urls of images")
	}
	//images will be downloaded and in compressed form
}
