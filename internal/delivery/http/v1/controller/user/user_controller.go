package userControllers

import (
	userHandler "beli-mang/internal/delivery/http/v1/controller/user/handler"
	userService "beli-mang/internal/service/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, userSvc userService.UserService, jwt echo.MiddlewareFunc) {
	user := group.Group("/users")
	handler := userHandler.NewHandler(userSvc, val)

	user.POST("/register", handler.ITRegister)
	user.POST("/login", handler.ITLogin)
}
