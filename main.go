package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"zocket.example.com/api"
	"zocket.example.com/consumer"
	"zocket.example.com/database"
	"zocket.example.com/producer"
)

func main() {

	//Connection to database
	DB, err := database.NewDatabase()
	if err != nil {
		log.Println("Cannot connect to database", err)
	}
	// Information for operating program

	log.Println("go to http://localhost:8002/createproduct")

	log.Println("example of Json Body", ` {
		"user_id":3,
	   "product_name":"Iphone",f
	   "product_description":"its a nice phone",
	   "product_images":["https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTzoXEEUkHp4Y3tRke237DccQoFTs4slk4JbA&usqp=CAU"],
	   "product_price":1000
	   }"`)
	// Consumer Routine
	go func() {
		err := consumer.ConsumeMessages(DB)
		if err != nil {
			log.Println(err)
		}
	}()
	//insatnce if producder
	rabbitMQProducer, err := producer.NewRabbitMQProducer("message_queue")
	if err != nil {
		log.Println("Error creating RabbitMQ producer:", err)
		return
	}
	defer rabbitMQProducer.Close()

	// Route for Creating New product by user, Instance of CreateProduct Handler
	newCreateProductInstance := api.CreateNewProductHandler(rabbitMQProducer, DB)
	http.HandleFunc("/createproduct", newCreateProductInstance)

	// start the server at 8002
	err = http.ListenAndServe(":8002", nil)

	if err != nil {
		log.Fatal("Error starting the HTTP server:", err)
	}

}
