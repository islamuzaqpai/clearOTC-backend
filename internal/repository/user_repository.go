package repository

import (
	"ClearOTC/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, userReq *models.UserReq) (*models.UserRes, error) {
	query := `INSERT INTO clearotc.users (name, lastname, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id, name, lastname, email, created_at, updated_at`

	user := &models.UserRes{}

	err := r.pool.QueryRow(
		ctx,
		query,
		userReq.Name,
		userReq.Lastname,
		userReq.Email,
		userReq.PasswordHash,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Lastname,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
