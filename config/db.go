package config

type (
	DbConfig struct {
		DbHost     string `envconfig:"DB_HOST" default:"localhost"`
		DbUser     string `envconfig:"DB_USER" default:"root"`
		DbName     string `envconfig:"DB_NAME" default:"local"`
		DbPassword string `envconfig:"DB_PASSWORD" default:"secret"`
		DbSSL      string `envconfig:"DB_SSL" default:"disable"`
		DbPort     string `envconfig:"DB_PORT" default:"5432"`
	}
)
