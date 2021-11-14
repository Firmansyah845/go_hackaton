package forecast

import (
	"context"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/dto"
)

type ServiceForecasting interface {
	GetDataForecasting(ctx context.Context, request []dto.LoginResponse) error
}
