package users

import "context"

type Repository interface {
	byID(ctx context.Context, id string) (User, error)
	byEmail(ctx context.Context, email string) (User, error)
	new(ctx context.Context, u User) (User, error)
	list(ctx context.Context) ([]User, error)
	remove(ctx context.Context, id string) error
}
