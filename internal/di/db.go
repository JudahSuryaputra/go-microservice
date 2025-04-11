package di

import (
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
