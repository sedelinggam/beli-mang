package userService

import (
	"beli-mang/internal/delivery/http/v1/request"
	"beli-mang/internal/delivery/http/v1/response"
	"beli-mang/internal/entity"
	"beli-mang/package/crypto/bcrypt"
	cryptoJWT "beli-mang/package/crypto/jwt"
	"beli-mang/package/lumen"
	"context"
	"errors"
	"strconv"
)

func (ss userService) Login(ctx context.Context, requestData request.UserLogin) (*response.UserAccessToken, error) {
	var (
		err error
	)
	//Check User NIP
	userData := entity.User{
		NIP:      strconv.Itoa(requestData.NIP),
		UserRole: requestData.RoleType,
	}

	//Check NIP
	if validNIP := userData.CheckNIP(true); !validNIP {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("NIP not valid"))
	} else if requestData.RoleType == "t" && userData.NIP[0:3] != "615" {
		return nil, lumen.NewError(lumen.ErrNotFound, errors.New("NIP not found"))
	} else if requestData.RoleType == "u" && userData.NIP[0:3] != "303" {
		return nil, lumen.NewError(lumen.ErrNotFound, errors.New("NIP not found"))
	}

	// Find the user by credentials
	user := entity.User{}

	//Compare password hash
	if !bcrypt.CheckPasswordHash(requestData.Password, user.Password) {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("password doesn't match"))
	}
	// Create the Claims
	accessToken, err := cryptoJWT.GenerateToken(user.ID, user.NIP, user.UserRole)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	respAccessToken := &response.UserAccessToken{
		UserID:      user.ID,
		NIP:         requestData.NIP,
		Name:        user.Name,
		AccessToken: *accessToken,
	}

	return respAccessToken, nil
}
