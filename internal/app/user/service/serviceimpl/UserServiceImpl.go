package serviceimpl

import (
	"context"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/dto"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/repository"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/service"
	"github.com/Firmansyah845/go_hackaton/internal/pkg/custom/forecast"
)

type UserServiceImpl struct {
	UserRepo        repository.UserRepository
	forecastService forecast.ServiceForecasting
}

func (r UserServiceImpl) GetData(c context.Context, userId int, fromDate, toDate string) (*[]dto.SalesResponse, error) {
	response, err := r.UserRepo.GetData(c, userId, fromDate, toDate)
	if err != nil {
		return response, err
	}

	//hit ke backend phyton

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
func CreateUserServiceImpl(repo repository.UserRepository, forecastService forecast.ServiceForecasting) service.UserService {
	return UserServiceImpl{repo, forecastService}
}
