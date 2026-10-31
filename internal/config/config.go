package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Host string
	Port int
	//ReadTimeout  time.Duration
	//WriteTimeout time.Duration
	//IdleTimeout  time.Duration
}

type Config struct {
	DatabaseCfg DatabaseConfig
	ServerCfg   ServerConfig
}

func LoadConfig(filename string) (*Config, error) {
	if err := godotenv.Load(filename); err != nil {
		return nil, fmt.Errorf("failed to load cfg: %w", err)
	}

	cfg := Config{}

	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse to int: %w", err)
	}

	cfg.DatabaseCfg.Port = dbPort
	cfg.DatabaseCfg.Host = os.Getenv("POSTGRES_HOST")
	cfg.DatabaseCfg.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.DatabaseCfg.User = os.Getenv("POSTGRES_USER")
	cfg.DatabaseCfg.Name = os.Getenv("POSTGRES_DB")

	cfg.ServerCfg.Host = os.Getenv("SERVER_HOST")

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse to int: %w", err)
	}
	cfg.ServerCfg.Port = serverPort

	err = ValidateCfg(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to validate cfg: %w", err)
	}

	return &cfg, nil
}

func ValidateCfg(cfg *Config) error {
	if cfg == nil {
		return fmt.Errorf("config is nil")
	}

	// Database validation
	if cfg.DatabaseCfg.Host == "" {
		return fmt.Errorf("POSTGRES_HOST is required")
	}
	if cfg.DatabaseCfg.Port <= 0 || cfg.DatabaseCfg.Port > 65535 {
		return fmt.Errorf("POSTGRES_PORT must be between 1 and 65535")
	}
	if cfg.DatabaseCfg.User == "" {
		return fmt.Errorf("POSTGRES_USER is required")
	}
	if cfg.DatabaseCfg.Password == "" {
		return fmt.Errorf("POSTGRES_PASSWORD is required")
	}
	if cfg.DatabaseCfg.Name == "" {
		return fmt.Errorf("POSTGRES_DB is required")
	}

	// Server validation
	if cfg.ServerCfg.Host == "" {
		return fmt.Errorf("SERVER_HOST is required")
	}
	if cfg.ServerCfg.Port <= 0 || cfg.ServerCfg.Port > 65535 {
		return fmt.Errorf("SERVER_PORT must be between 1 and 65535")
	}

	return nil
}
