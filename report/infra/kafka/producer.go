package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaProducer() KafkaProducer {
	return KafkaProducer{}
}

func (kafka *KafkaProducer) SetupProducer(bootstrapServers string) {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	}

	kafka.Producer, _ = ckafka.NewProducer(configMap)
}

func (kafka *KafkaProducer) Publish(msg string, topic string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := kafka.Producer.Produce(message, kafka.DeliveryChan)
	if err != nil {
		return err
	}

	return nil
}

func (kafka *KafkaProducer) DeliveryReport() {
	for e := range kafka.DeliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}
}
