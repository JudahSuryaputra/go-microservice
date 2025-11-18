package config

type JwtConfig struct {
	JwtSecret string `envconfig:"JWT_SECRET" default:"secret"`
	JwtTtl    string `envconfig:"JWT_TTL" default:"3600"`
}
