package consumer

import (
	"fmt"
	"reflect"
	"testing"

	"zocket.example.com/database"
)

func TestConsumerFunction(t *testing.T) {

	DB, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Cannot connect to database", err)
	}
	err = ConsumerFunction("1", DB)
	if err != nil {
		t.Errorf("Consumer Function Failed")
	}

}

func TestDownloadImage(t *testing.T) {
	imageURL := "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTzoXEEUkHp4Y3tRke237DccQoFTs4slk4JbA&usqp=CAU"
	res, err := DownloadImage(imageURL)
	if err != nil {
		t.Errorf("Cannot get image data: %v", err)
		return
	}

	expectedType := reflect.TypeOf([]uint8(nil))
	resultType := reflect.TypeOf(res)

	if resultType != expectedType {
		t.Errorf("Expected type %v, but got type %v", expectedType, resultType)
	}

}
