package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	// Env  string `env:"ENV" envDefault:"dev"`
	Port       string `env:"PORT" envDefault:"8080"`
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
}

// func LoadConfig(envfile ...string) *Config {
// 	cfg := &Config{}

// 	_ = godotenv.Load(envfile...)
// 	_ = env.Parse(cfg)

// 	return cfg
// }

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load()

	if err != nil {
		return cfg, err
	}

	err = env.Parse(cfg)

	return cfg, err
}

func NewDBConnection(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
