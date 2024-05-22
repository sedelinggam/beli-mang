package userService

import (
	"beli-mang/internal/delivery/http/v1/request"
	"beli-mang/internal/delivery/http/v1/response"
	userRepository "beli-mang/internal/repository/user"
	"context"

	"github.com/jmoiron/sqlx"
)

type userService struct {
	userRepo userRepository.UserRepository
}

type UserService interface {
	Login(ctx context.Context, requestData request.UserLogin) (*response.UserAccessToken, error)
	RegisterUserIT(ctx context.Context, requestData request.ITUserRegister) (*response.UserAccessToken, error)
	GetUsers(ctx context.Context, req request.UserParam) ([]*response.UserNurse, error)
}

func New(db *sqlx.DB) UserService {
	return &userService{
		userRepo: userRepository.New(db),
	}
}
