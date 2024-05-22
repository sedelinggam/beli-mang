package userRepository

import (
	"beli-mang/internal/delivery/http/v1/request"
	"beli-mang/internal/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	Create(ctx context.Context, data entity.User) error
	GetUsers(ctx context.Context, req request.UserParam) ([]*entity.User, error)
}

func New(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
