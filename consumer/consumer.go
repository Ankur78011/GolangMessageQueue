package consumer

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/streadway/amqp"
	"zocket.example.com/database"
)

func ConsumeMessages(Db *database.Database) error {

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	defer connection.Close()
	log.Println("Successfully connected to RabbitMQ instance")
	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer channel.Close()
	queue, err := channel.QueueDeclare(
		"message_queue", // name
		false,           // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // args
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}
	messages, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return fmt.Errorf("failed to consume messages: %v", err)
	}

	go func() {
		for message := range messages {
			log.Printf("Received a message: %s\n", message.Body)
			err = ConsumerFunction(string(message.Body), Db)
			if err != nil {
				fmt.Printf("Error=%s", err)
			}
		}
	}()

	log.Println("Consumer is waiting for messages. To exit, press Ctrl+C")
	select {}
}

func ConsumerFunction(productID string, Db *database.Database) error {
	productIDInt, _ := strconv.Atoi(productID)
	var compressedImagesLocalPath []string
	// Getting the path of current directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Print("cannot get local path", err)
		return err
	}
	arrayOfImageUrl, err := Db.GetProductFromProductId(productIDInt)
	if err != nil {
		fmt.Print("cannot fetch iamge URls", err)
		return err
	}

	for i, imageURL := range arrayOfImageUrl {
		imageData, err := DownloadImage(imageURL)
		if err != nil {
			fmt.Println(" Error downloading image:", err)
		}
		filename := fmt.Sprintf("compressed_image_%s_%s.jpg.gz", strconv.Itoa(i), productID)
		compressedLocalFile := currentDir + filename
		compressedImagesLocalPath = append(compressedImagesLocalPath, compressedLocalFile)

		err = CompressAndSaveImage(imageData, filename)
		if err != nil {
			fmt.Println("Error compressing and saving image:", err)
		}
	}
	err = Db.AddCompressedProductImages(productIDInt, compressedImagesLocalPath)
	if err != nil {
		fmt.Println("Cannot send compressed files to database ", err)
	}
	log.Println("Images downloaded, compressed, and saved successfully and local path stored in the database")
	return nil

}

func DownloadImage(url string) ([]byte, error) {
	// Send HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the image data
	imageData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}

func CompressAndSaveImage(imageData []byte, filename string) error {
	// Create a new compressed file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	// Write the image data to the gzip writer
	_, err = gzipWriter.Write(imageData)
	if err != nil {
		return err
	}

	return nil
}
