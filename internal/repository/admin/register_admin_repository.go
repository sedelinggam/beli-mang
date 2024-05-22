package adminRepository

import (
	"beli-mang/internal/entity"
	"context"
	"fmt"
)

func (ar adminRepository) Register(ctx context.Context, data entity.Admin) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, username, password, email, created_at) VALUES (:id, :username, :password, :email, :created_at)`, data.TableName())
	_, err := ar.db.NamedExec(query, data)
	if err != nil {
		return err
	}
	return nil
}
