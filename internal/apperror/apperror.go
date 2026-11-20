package apperror

import "errors"

var (
	ErrRequiredField      = errors.New("required field is empty")
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidInput       = errors.New("invalid input")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrEmailAlreadyExists = errors.New("email already exists")
)
