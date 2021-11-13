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

// GetData repository instance
func (b UserRepoImpl) GetData(c context.Context, userId int, fromDate, toDate string) (*[]dto.SalesResponse, error) {

	var (
		response []dto.SalesResponse
		each     dto.SalesResponse
	)

	param := "BETWEEN '" + fromDate + "' AND '" + toDate + "'"

	args := fmt.Sprintf(`SELECT id, sales_date, value, user_id FROM sales WHERE (sales_date %s) AND user_id = %d`, param, userId)

	if rows, err := config.App.DB.QueryContext(c, args); err != nil {
		logger.WithFields(logger.Fields{}).Errorf("error get data sales: " + err.Error())
		return nil, fmt.Errorf("error get data sales : " + err.Error())
	} else {
		defer rows.Close()
		for rows.Next() {

			if err = rows.Scan(&each.ID, &each.SalesDate, &each.Value, &each.UserId); err != nil {
				logger.WithFields(logger.Fields{}).Infof("ERROR: Can't reach data sales :" + err.Error())
			}
			response = append(response, each)
		}
	}

	return &response, nil
}

// CreateUserRepoImpl create user repository instance
func CreateUserRepoImpl() repository.UserRepository {
	return &UserRepoImpl{}
}
