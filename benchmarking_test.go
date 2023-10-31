package main

import (
	"net/http"
	"testing"

	"zocket.example.com/consumer"
	"zocket.example.com/database"
)

// benchmarked  HTTP request to create a product
func BenchmarkHTTPCreateProduct(b *testing.B) {

	for i := 0; i < b.N; i++ {

		req, err := http.NewRequest("POST", "http://localhost:8002/createproduct", nil)
		if err != nil {
			b.Fatalf("Error creating request: %v", err)
		}

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			b.Fatalf("Error making request: %v", err)
		}
	}
}

// becnhmark for download and compress
func BenchmarkImageDownloadAndCompression(b *testing.B) {

	for i := 0; i < b.N; i++ {

		res, err := consumer.DownloadImage("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR9lnaHnIsJ5QUoVgwts4eSJibGWYc309Nugw&usqp=CAU:")
		if err != nil {
			b.Fatalf("Error downloading and compressing images: %v", err)
		}

		for i := 0; i < b.N; i++ {

			err := consumer.CompressAndSaveImage(res, "compressed_image_")
			if err != nil {
				b.Fatalf("Error downloading and compressing images: %v", err)
			}
		}

	}

}

// becnhmark for
func BenchmarkConsumerFunction(b *testing.B) {
	DB, err := database.NewDatabase()
	if err != nil {
		b.Fatalf("Error in database connection: %v", err)
	}
	for i := 0; i < b.N; i++ {
		err := consumer.ConsumerFunction("88", DB)
		if err != nil {
			b.Fatalf("Error in consumer function: %v", err)

		}
	}
}
