package main

import (
	"ClearOTC/internal/config"
	"ClearOTC/internal/database"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		return fmt.Errorf("failed to load cfg: %w", err)
	}

	pool, err := database.Connect(&cfg.DatabaseCfg)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}

	defer pool.Close()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.ServerCfg.Host, cfg.ServerCfg.Port),
		Handler: http.DefaultServeMux,
	}

	serverErr := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("shutdown signal received")
	case err = <-serverErr:
		return fmt.Errorf("http server failed: %w", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}

	return nil
}
