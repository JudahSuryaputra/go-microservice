package kafka

import (
	"github.com/Shopify/sarama"
	"log"
)

func InitKafkaProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	return producer
}
