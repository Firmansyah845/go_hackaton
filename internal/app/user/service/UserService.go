package service

import (
	"context"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/dto"
)

type UserService interface {
	Login(c context.Context, username, password string) (*dto.LoginResponse, error)
}
