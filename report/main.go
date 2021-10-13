package main

import (
	"fmt"
	"log"
	"os"

	"github.com/HRsniper/imersao-fullstack-fullcycle-4/report/infra/kafka"
	"github.com/HRsniper/imersao-fullstack-fullcycle-4/report/infra/repository"
	"github.com/HRsniper/imersao-fullstack-fullcycle-4/report/usecase"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

func main() {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTIC_SEARCH_CLIENT_ADDRESS")},
	})
	if err != nil {
		log.Fatalf("error connecting to elasticsearch")
	}

	repo := repository.TransactionElasticRepository{
		Client: *client,
	}

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		err := usecase.GenerateReport(msg.Value, repo)
		if err != nil {
			fmt.Println(err)
		}
	}
}
