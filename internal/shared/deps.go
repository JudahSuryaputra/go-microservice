package shared

import (
	"github.com/redis/go-redis/v9"
	"go-microservice/config"
	kafka2 "go-microservice/internal/di/kafka"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type (
	Deps struct {
		dig.In

		DB          *gorm.DB
		Config      *config.Configuration
		RedisClient *redis.Client
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(kafka2.InitKafkaConsumer); err != nil {
		panic(err)
	}
	if err := container.Provide(kafka2.InitKafkaProducer); err != nil {
		panic(err)
	}
	return nil
}
