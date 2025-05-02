package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID    string `db:"id"`
	Email string `db:"email"`
	Name  string `db:"name"`
}

type UserService struct {
	db *DB
}

func NewUserService(db *DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) UserByID(ctx context.Context, id string) (*User, error) {
	query := `
	SELECT
		users.id,
		users.email,
		users.name
	FROM users
	WHERE users.id = $1`
	rows, err := s.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func (s *UserService) CreateUser(ctx context.Context, id, email, name string) error {
	_, err := s.db.Exec(ctx, `
		INSERT INTO users (id, email, name)
		VALUES ($1, $2, $3)
	`, id, email, name)
	return err
}

func (s *UserService) UpdateUser(ctx context.Context, id, email, name string) error {
	_, err := s.db.Exec(ctx, `
		UPDATE users
		SET email = $2,
			name = $3,
			updated_at = (now() at time zone 'utc')
		WHERE id = $1
	`, id, email, name)
	return err
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	_, err := s.db.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	return err
}
