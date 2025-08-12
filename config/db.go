package config

type (
	DbConfig struct {
		DbHost     string `envconfig:"DB_HOST" required:"true" default:"localhost"`
		DbUser     string `envconfig:"DB_USER" required:"true" default:"postgres"`
		DbName     string `envconfig:"DB_NAME" required:"true" default:"postgres"`
		DbPassword string `envconfig:"DB_PASSWORD" required:"true" default:"postgres"`
		DbSSL      string `envconfig:"DB_SSL" required:"true" default:"disable"`
		DbPort     string `envconfig:"DB_PORT" required:"true" default:"5432"`
	}
)
