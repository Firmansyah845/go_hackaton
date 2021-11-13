package repository

import (
	"context"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/dto"
)

type UserRepository interface {
	Login(c context.Context, username, password string) (*dto.LoginResponse, error)
	GetData(c context.Context, userId int, fromDate, toDate string) (*[]dto.SalesResponse, error)
}
