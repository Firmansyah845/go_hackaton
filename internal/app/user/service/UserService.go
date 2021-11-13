package service

import (
	"context"
	"go-hackaton/internal/app/user/dto"
)

type UserService interface {
	Login(c context.Context, username, password string) (*dto.LoginResponse, error)
}
