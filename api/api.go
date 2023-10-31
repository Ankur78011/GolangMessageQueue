package api

import (
	"encoding/json"
	"log"
	"net/http"

	"zocket.example.com/database"
	"zocket.example.com/models"
	"zocket.example.com/producer"
)

func CreateNewProductHandler(producer *producer.RabbitMQProducer, Db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProduct models.NewProduct
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Invalid Detials", http.StatusBadRequest)
			return
		}
		//Get the Id of product
		productID, err := Db.StoreProduct(&newProduct)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
		}
		//passing product_id to the producer
		err = producer.PublishMessage(int(productID))
		if err != nil {
			log.Println("Cannot send Id to Producer", err)
			http.Error(w, "Invalid Input", http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
	}
}
