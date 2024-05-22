package adminService

import (
	"beli-mang/internal/delivery/http/v1/request"
	"beli-mang/internal/delivery/http/v1/response"
	"beli-mang/internal/entity"
	"beli-mang/package/crypto/bcrypt"
	cryptoJWT "beli-mang/package/crypto/jwt"
	"beli-mang/package/lumen"
	"context"
	"time"

	"github.com/oklog/ulid/v2"
)

func (as adminService) Register(ctx context.Context, req request.AdminRegister) (*response.AdminRegister, error) {
	var (
		err          error
		hashPassword string
	)

	hashPassword, err = bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	adminData := entity.Admin{
		ID:        ulid.Make().String(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}

	_, err = as.adminRepo.GetAdminByUsername(ctx, adminData.Username)

	if err != nil {
		if lumen.CheckErrorSQLUnique(err) {
			return nil, lumen.NewError(lumen.ErrConflict, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	token, err := cryptoJWT.GenerateToken(adminData.ID, adminData.Username, adminData.Email)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	respToken := response.AdminRegister{
		Token: *token,
	}

	return &respToken, nil
}
