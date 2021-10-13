package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumer holds all consumer logic and settings of Apache Kafka connections/
// Also has a Message channel which is a channel where the messages are going to be pushed
type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

// NewKafkaConsumer creates a new KafkaConsumer struct with its message channel as dependency
func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

// Consume consumes all message pulled from apache kafka and sent it to message channel
func (kafka *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":          os.Getenv("KAFKA_CONSUMER_GROUP_ID"),
		// "security.protocol": os.Getenv("security.protocol"),
		// "sasl.mechanisms":   os.Getenv("sasl.mechanisms"),
		// "sasl.username":     os.Getenv("sasl.username"),
		// "sasl.password":     os.Getenv("sasl.password"),
	}

	consumer, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message:" + err.Error())
	}

	topics := []string{os.Getenv("KAFKA_READ_TOPIC")}
	consumer.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			kafka.MsgChan <- msg
		}
	}
}
