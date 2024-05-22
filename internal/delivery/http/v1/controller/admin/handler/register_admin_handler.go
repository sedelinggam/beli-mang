package adminHandler

import (
	"beli-mang/internal/delivery/http/v1/request"
	"beli-mang/internal/delivery/http/v1/response"
	"beli-mang/package/lumen"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ah adminHandler) Register(c echo.Context) error {
	var (
		req  request.AdminRegister
		resp *response.AdminRegister
		err  error
	)
	err = c.Bind(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	// Validate the admin struct
	err = ah.val.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	resp, err = ah.adminService.Register(c.Request().Context(), req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusCreated, response.Common{
		Message: "admin registered successfully",
		Data:    resp,
	})
}
