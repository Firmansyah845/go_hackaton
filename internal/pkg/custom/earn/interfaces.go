package earn

import (
	"context"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/dto"
)

type ServiceMonetize interface {
	CreateEarn(ctx context.Context, request []dto.LoginResponse) error
}
