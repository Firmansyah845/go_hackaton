package serviceimpl

import (
	"context"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/dto"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/repository"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/service"
	"github.com/Firmansyah845/go_hackaton/internal/pkg/custom/earn"
)

type UserServiceImpl struct {
	UserRepo        repository.UserRepository
	MonetizeService earn.ServiceMonetize
}

func (r UserServiceImpl) GetData(c context.Context, userId int, fromDate, toDate string) (*[]dto.SalesResponse, error) {
	response, err := r.UserRepo.GetData(c, userId, fromDate, toDate)
	if err != nil {
		return response, err
	}

	return response, nil
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
