package main

import (
	"ClearOTC/internal/config"
	"ClearOTC/internal/database"
	"fmt"
)

func Run() error {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		return fmt.Errorf("failed to load cfg: %w", err)
	}

	pool, err := database.Connect(&cfg.DatabaseCfg)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %v", err)
	}

	fmt.Println(pool)
	return nil
}
