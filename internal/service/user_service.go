package service

import (
	"ClearOTC/internal/models"
	"ClearOTC/internal/utils"
	"context"
	"fmt"
	"strings"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
}

type UserService struct {
	repository UserRepository
}

func NewService(repository UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Register(ctx context.Context, req *models.CreateUserParams) (*models.UserRes, error) {
	if err := utils.ValidateRegister(req); err != nil {
		return nil, fmt.Errorf("validate registration: %w", err)
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user := &models.User{
		Name:         strings.TrimSpace(req.Name),
		Lastname:     strings.TrimSpace(req.Lastname),
		Email:        strings.TrimSpace(req.Email),
		PasswordHash: hashedPassword,
	}

	if err := s.repository.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &models.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
