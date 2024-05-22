package userHandler

import (
	userService "beli-mang/internal/service/user"

	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService userService.UserService
	val         *validator.Validate
}

func NewHandler(userService userService.UserService, val *validator.Validate) *userHandler {
	return &userHandler{
		userService: userService,
		val:         val,
	}
}
