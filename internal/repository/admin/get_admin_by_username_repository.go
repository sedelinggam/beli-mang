package adminRepository

import (
	"beli-mang/internal/entity"
	"context"
)

func (ar adminRepository) GetAdminByUsername(ctx context.Context, username string) (*entity.Admin, error) {
	var (
		data entity.Admin
	)
	query := "SELECT * FROM admins WHERE username = $1"
	err := ar.db.Get(&data, query, username)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
