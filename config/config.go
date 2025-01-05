package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Server ServerConfig
	DB     DatabaseConfig
	Debug  bool `env:"DEBUG,required"`
}

type ServerConfig struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

type DatabaseUser struct {
	Name     string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
}

type DatabaseConfig struct {
	Host string `env:"DB_HOST,required"`
	Port int    `env:"DB_PORT,required"`
	User DatabaseUser
	Name string `env:"DB_NAME,required"`
}

func LoadConfig() *Config {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("error: environment: failed to load environment %s", err)
	}

	return &cfg
}

func LoadDatabaseConfig() *DatabaseConfig {
	var cfg DatabaseConfig

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("error: environment: failed to load environment %s", err)
	}

	return &cfg
}
