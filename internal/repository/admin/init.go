package adminRepository

import (
	"beli-mang/internal/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type adminRepository struct {
	db *sqlx.DB
}

type AdminRepository interface {
	Register(ctx context.Context, data entity.Admin) error
	GetAdminByUsername(ctx context.Context, username string) (*entity.Admin, error)
}

func New(db *sqlx.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}
