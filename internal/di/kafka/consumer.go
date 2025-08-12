package kafka

import (
	"github.com/Shopify/sarama"
	"log"
)

func InitKafkaConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	return consumer
}
