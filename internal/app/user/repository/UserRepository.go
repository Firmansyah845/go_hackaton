package repository

import (
	"context"
	"go-hackaton/internal/app/user/dto"
)

type UserRepository interface {
	Login(c context.Context, username, password string) (*dto.LoginResponse, error)
}
