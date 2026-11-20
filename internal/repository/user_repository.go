package repository

import (
	"ClearOTC/internal/apperror"
	"ClearOTC/internal/models"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO clearotc.users (name, lastname, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	err := r.pool.QueryRow(
		ctx,
		query,
		user.Name,
		user.Lastname,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) &&
			pgErr.Code == "23505" &&
			pgErr.ConstraintName == "users_email_unique_idx" {
			return apperror.ErrEmailAlreadyExists
		}

		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
