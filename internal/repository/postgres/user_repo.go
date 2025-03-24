package postgres

import (
	"context"
	"go-automation/internal/domain/user"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*user.User, error) {
	// запрос к базе
}

func (r *UserRepo) Create(ctx context.Context, u *user.User) error {
	// insert в базу
}
