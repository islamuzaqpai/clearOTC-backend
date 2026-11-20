package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Lastname     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type CreateUserParams struct {
	Name         string
	Lastname     string
	Email        string
	PasswordHash string
}

type UserRes struct {
	ID        uuid.UUID
	Name      string
	Lastname  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
