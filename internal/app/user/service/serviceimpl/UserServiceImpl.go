package serviceimpl

import (
	"context"
	"go-hackaton/internal/app/user/dto"
	"go-hackaton/internal/app/user/repository"
	"go-hackaton/internal/app/user/service"
	"go-hackaton/internal/pkg/custom/earn"
)

type UserServiceImpl struct {
	UserRepo       repository.UserRepository
	MonetizeService earn.ServiceMonetize
}

func (r UserServiceImpl) Login(c context.Context, username, password string) (*dto.LoginResponse, error) {

	response, err := r.UserRepo.Login(c, username, password)
	if err != nil {
		return response, err
	}

	return response, nil
}

// CreateUserServiceImpl create user service instance
func CreateUserServiceImpl(repo repository.UserRepository, monetizeService earn.ServiceMonetize) service.UserService {
	return UserServiceImpl{repo, monetizeService}
}
