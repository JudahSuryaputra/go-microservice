package di

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-microservice/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func NewORM(cfg *config.Configuration) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		cfg.DbHost,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbName,
		cfg.DbPort,
		cfg.DbSSL,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // change to Info for dev
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Get the underlying *sql.DB to configure pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get DB from GORM: %v", err)
	}

	// Pooling config
	sqlDB.SetMaxIdleConns(5)                   // idle connections kept alive
	sqlDB.SetMaxOpenConns(20)                  // max connections in pool
	sqlDB.SetConnMaxLifetime(time.Hour)        // recycle connections hourly
	sqlDB.SetConnMaxIdleTime(30 * time.Minute) // max idle time

	// Ping DB to verify connectivity
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping database: %v", err)
	}

	log.Println("✅ Database connected with connection pooling")
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
