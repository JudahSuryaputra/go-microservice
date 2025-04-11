package di

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-microservice/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewORM(cfg *config.Configuration) *gorm.DB {
	dsn := "host=" + cfg.DbHost + " user=" + cfg.DbUser + " password=" + cfg.DbPassword +
		" dbname=" + cfg.DbName + " port=" + cfg.DbPort + " sslmode=" + cfg.DbSSL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func NewRedisClient(cfg *config.Configuration) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisServer,
	})
	result, err := client.Ping(context.Background()).Result()
	fmt.Printf("Redis ping result: %v\n", result)
	if err != nil {
		panic(err)
	}
	return client
}
