package shared

import (
	"go-microservice/config"
	"go-microservice/internal/kafka"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type (
	Deps struct {
		dig.In
		DB     *gorm.DB
		Config *config.Configuration
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(kafka.InitKafkaConsumer); err != nil {
		panic(err)
	}
	if err := container.Provide(kafka.InitKafkaProducer); err != nil {
		panic(err)
	}

	return nil
}
