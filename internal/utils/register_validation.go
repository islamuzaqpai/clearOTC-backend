package utils

import (
	"ClearOTC/internal/apperror"
	"ClearOTC/internal/models"
	"net/mail"
	"strings"
	"unicode/utf8"
)

func ValidateRegister(req *models.CreateUserParams) error {
	if req == nil {
		return apperror.ErrInvalidInput
	}

	name := strings.TrimSpace(req.Name)
	lastname := strings.TrimSpace(req.Lastname)
	email := strings.TrimSpace(req.Email)

	// Required fields
	if name == "" || lastname == "" || email == "" ||
		strings.TrimSpace(req.Password) == "" {
		return apperror.ErrRequiredField
	}

	// Name and lastname length
	if utf8.RuneCountInString(name) < 2 ||
		utf8.RuneCountInString(name) > 100 ||
		utf8.RuneCountInString(lastname) < 2 ||
		utf8.RuneCountInString(lastname) > 100 {
		return apperror.ErrInvalidInput
	}

	// Email validation
	addr, err := mail.ParseAddress(email)
	if err != nil || addr.Address != email {
		return apperror.ErrInvalidEmail
	}

	// Password validation
	if len(req.Password) < 8 || len(req.Password) > 72 {
		return apperror.ErrInvalidPassword
	}

	return nil
}
