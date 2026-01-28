package users

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, email string, name string) (User, error) {
	if err := ValidateNewUser(email, name); err != nil {
		return User{}, err
	}

	if _, err := s.repo.GetByEmail(ctx, email); err == nil {
		return User{}, ErrEmailTaken
	}

	u := User{
		ID:    uuid.NewString(),
		Email: email,
		Name:  name,
	}

	u, err := s.repo.Create(ctx, u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *Service) Get(ctx context.Context, id string) (User, error) {
	return s.repo.GetByID(ctx, id)
}
