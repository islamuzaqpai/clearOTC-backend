package database

import (
	"ClearOTC/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(dbCfg *config.DatabaseConfig) (*pgxpool.Pool, error) {
	if dbCfg == nil {
		return nil, fmt.Errorf("database config is nil")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s",
		dbCfg.Host,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Port,
		dbCfg.Name,
	)

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse to pool cfg: %w", err)
	}

	poolCfg.MinConns = 2
	poolCfg.HealthCheckPeriod = time.Minute * 5
	poolCfg.MaxConns = 10

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
