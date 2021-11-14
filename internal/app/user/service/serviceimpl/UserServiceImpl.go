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

func (r UserServiceImpl) GetData(c context.Context, userId, period int, fromDate, toDate, fromDateActual string) (*dto.ResponseForecast, *[]dto.SalesResponse, error) {
	_, data, err := r.UserRepo.GetData(c, userId, fromDate, toDate)
	if err != nil {
		return nil, nil, err
	}

	//hit ke backend phyton
	request := dto.PayloadForecast{
		UserId: userId,
		Period: period,
		Data:   *data,
	}
	responseForecast, err := r.forecastService.GetDataForecasting(c, request)
	if err != nil {
		return nil, nil, err
	}

	responseActual, _, err := r.UserRepo.GetData(c, userId, fromDateActual, toDate)
	if err != nil {
		return nil, nil, err
	}

	return responseForecast, responseActual, nil
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
