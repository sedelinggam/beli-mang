package adminHandler

import (
	adminService "beli-mang/internal/service/admin"

	"github.com/go-playground/validator/v10"
)

type adminHandler struct {
	adminService adminService.AdminService
	val          *validator.Validate
}

func NewHandler(adminService adminService.AdminService, val *validator.Validate) *adminHandler {
	return &adminHandler{
		adminService: adminService,
		val:          val,
	}
}
