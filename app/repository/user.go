package repository

import (
	"fmt"
	"errors"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nekto-sns/nekto-server/app/model"
)

type userRepository struct{
	db *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) (*userRepository) {
	return &userRepository{
		db: pool,
	}
}

func (r *userRepository) ByName(ctx context.Context, name string) (*model.User, error) {
	var u model.User

	query := "SELECT id, name, display_name, bio, created_at FROM users WHERE name = $1"
	err := r.db.QueryRow(ctx, query, name).Scan(&u.ID, &u.Name, &u.DisplayName, &u.Bio, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("repository: Row not found: %w", model.ErrNotFound)
		}
		return nil, fmt.Errorf("repository: DB query failed for username=%s : %w (%v)", name, model.ErrInternal, err)
	}

	return &u, nil
}
