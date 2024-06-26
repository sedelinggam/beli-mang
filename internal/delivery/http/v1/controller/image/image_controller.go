package imageController

import (
	imageHandler "beli-mang/internal/delivery/http/v1/controller/image/handler"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate, jwt echo.MiddlewareFunc) {
	image := group.Group("/image")
	handler := imageHandler.NewHandler(val)

	privateRoute := image
	privateRoute.Use(jwt)
	// TODO: Add middleware
	privateRoute.POST("", handler.StoreImage)
}
