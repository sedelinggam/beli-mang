package userControllers

import (
	adminHandler "beli-mang/internal/delivery/http/v1/controller/admin/handler"
	adminService "beli-mang/internal/service/admin"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, adminSvc adminService.AdminService, jwt echo.MiddlewareFunc) {
	user := group.Group("/admin")
	handler := adminHandler.NewHandler(adminSvc, val)

	user.POST("/register", handler.Register)
	user.POST("/login", handler.Register)
}
