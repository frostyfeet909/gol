package users

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	db *pgxpool.Pool
}

func NewPostgresRepo(db *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) GetByID(ctx context.Context, id string) (User, error) {
	const q = `
SELECT id, email, name
FROM users
WHERE id = $1
`
	var u User
	err := r.db.QueryRow(ctx, q, id).Scan(&u.ID, &u.Email, &u.Name)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, ErrNotFound
	}
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (r *PostgresRepo) GetByEmail(ctx context.Context, email string) (User, error) {
	const q = `
SELECT id, email, name
FROM users
WHERE email = $1
`
	var u User
	err := r.db.QueryRow(ctx, q, email).Scan(&u.ID, &u.Email, &u.Name)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, ErrNotFound
	}
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (r *PostgresRepo) Create(ctx context.Context, u User) (User, error) {
	const q = `
INSERT INTO users (id, email, name)
VALUES ($1, $2, $3)
`
	_, err := r.db.Exec(ctx, q, u.ID, u.Email, u.Name)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
