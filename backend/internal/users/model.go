package users

import (
	"errors"
	"strings"
)

var (
	ErrNotFound     = errors.New("user not found")
	ErrEmailTaken   = errors.New("email already taken")
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidName  = errors.New("invalid name")
)

type User struct {
	ID    string
	Email string
	Name  string
}

func ValidateNewUser(email string, name string) error {
	email = strings.TrimSpace(email)
	name = strings.TrimSpace(name)

	if len(name) < 2 {
		return ErrInvalidName
	}
	if !strings.Contains(email, "@") || len(email) < 5 {
		return ErrInvalidEmail
	}

	return nil
}
