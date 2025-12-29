package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	// Env  string `env:"ENV" envDefault:"dev"`
	Port string `env:"PORT" envDefault:"8080"`
}

// func LoadConfig(envfile ...string) *Config {
// 	cfg := &Config{}

// 	_ = godotenv.Load(envfile...)
// 	_ = env.Parse(cfg)

// 	return cfg
// }

func NewConfig() (Config, error) {
	cfg := Config{}

	err := godotenv.Load()

	if err != nil {
		return cfg, err
	}

	err = env.Parse(&cfg)

	return cfg, err
}
