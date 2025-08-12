package config

type RateLimitConfig struct {
	RPS   int `envconfig:"RPS" default:"5"`
	Burst int `envconfig:"BURST" default:"10"`
}
