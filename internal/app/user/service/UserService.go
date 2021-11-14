package service

import (
	"context"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/dto"
)

type UserService interface {
	Login(c context.Context, username, password string) (*dto.LoginResponse, error)
	GetData(c context.Context, userId, period int, fromDate, toDate, fromDateActual string) (*dto.ResponseForecast, *[]dto.SalesResponse, error)
}
