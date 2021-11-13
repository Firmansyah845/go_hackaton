package earn

import (
	"context"
	"go-hackaton/internal/app/user/dto"
)

type ServiceMonetize interface {
	CreateEarn(ctx context.Context, request []dto.LoginResponse) error
}
