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

func (s *Service) create(ctx context.Context, email string, name string) (User, error) {
	if _, err := s.repo.byEmail(ctx, email); err == nil {
		return User{}, ErrEmailTaken
	}

	u := User{
		ID:    uuid.NewString(),
		Email: email,
		Name:  name,
	}

	u, err := s.repo.new(ctx, u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *Service) get(ctx context.Context, id string) (User, error) {
	return s.repo.byID(ctx, id)
}

func (s *Service) list(ctx context.Context) ([]User, error) {
	return s.repo.list(ctx)
}

func (s *Service) remove(ctx context.Context, id string) error {
	return s.repo.remove(ctx, id)
}
