package service

import (
	"ClearOTC/internal/models"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, userReq *models.UserReq) (*models.UserRes, error)
}
