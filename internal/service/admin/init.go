package adminService

import (
	"beli-mang/internal/delivery/http/v1/request"
	"beli-mang/internal/delivery/http/v1/response"
	adminRepository "beli-mang/internal/repository/admin"
	"context"

	"github.com/jmoiron/sqlx"
)

type adminService struct {
	adminRepo adminRepository.AdminRepository
}

type AdminService interface {
	Register(ctx context.Context, data request.AdminRegister) (*response.AdminRegister, error)
}

func New(db *sqlx.DB) AdminService {
	return &adminService{
		adminRepo: adminRepository.New(db),
	}
}
