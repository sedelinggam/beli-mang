package userService

import (
	"context"
	"beli-mang/internal/delivery/http/v1/request"
	"beli-mang/internal/delivery/http/v1/response"
	"beli-mang/package/lumen"
)

func (ps userService) GetUsers(ctx context.Context, req request.UserParam) ([]*response.UserNurse, error) {
	products, err := ps.userRepo.GetUsers(ctx, req)
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	productsResp := response.MapUserListEntityToListResponse(products)
	return productsResp, nil
}
