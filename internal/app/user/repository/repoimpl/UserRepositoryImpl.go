package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Firmansyah845/go_hackaton/config"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/dto"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/repository"
	"github.com/Firmansyah845/go_hackaton/utils/logger"
)

// UserRepoImpl dependency
type UserRepoImpl struct {
}

// Login repository instance
func (b UserRepoImpl) Login(c context.Context, username, password string) (*dto.LoginResponse, error) {

	var (
		response     dto.LoginResponse
		passwordTemp sql.NullString
	)

	args := fmt.Sprintf(`SELECT id, name, username, password, role FROM user_data WHERE username = '%s'`, username)

	if err := config.App.DB.QueryRowContext(c, args).Scan(&response.UserId, &response.Name, &response.Username, &passwordTemp, &response.Role); err != nil {
		logger.WithFields(logger.Fields{}).Errorf("user account not exist : " + err.Error())
		return nil, fmt.Errorf("user account not exist")
	}

	if password != passwordTemp.String {
		return nil, fmt.Errorf("password not match")
	}

	return &response, nil
}

// CreateUserRepoImpl create user repository instance
func CreateUserRepoImpl() repository.UserRepository {
	return &UserRepoImpl{}
}
