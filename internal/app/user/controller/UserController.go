package controller

import (
	"github.com/Firmansyah845/go_hackaton/internal/app/user/service"
	"github.com/Firmansyah845/go_hackaton/utils/helper"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"time"
)

type UserController struct {
	userService service.UserService
}

func CreateUserController(router *echo.Echo, userService service.UserService) {
	inDB := UserController{userService: userService}

	api := router.Group("/v1")
	{
		userApi := api.Group("/hackaton")
		userApi.POST("/login", inDB.login)
		userApi.GET("/health/check", inDB.Health)
		userApi.POST("/data-filter-salesman", inDB.validate)
		userApi.POST("/data-default-salesman", inDB.DataDefault)
	}

}

func (b *UserController) login(c echo.Context) error {

	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}

	validate := helper.ValidateRequestFormData(c, rules)
	if validate != nil {
		return helper.ReturnResponse(
			c,
			http.StatusUnprocessableEntity,
			validate,
			nil,
			"validation error",
			"validation error")
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	response, err := b.userService.Login(c.Request().Context(), username, password)
	if err != nil {
		return helper.ReturnInvalidResponse(http.StatusBadRequest, nil, nil, err.Error(), err.Error())
	}

	return helper.ReturnResponse(c, http.StatusOK, response, nil,
		"Success", "Success")

}

func (b *UserController) validate(c echo.Context) error {

	rules := govalidator.MapData{
		"from_date": []string{"required"},
		"to_date":   []string{"required"},
		"interval":  []string{"required", "numeric"},
		"user_id":   []string{"required", "numeric"},
	}

	validate := helper.ValidateRequestFormData(c, rules)
	if validate != nil {
		return helper.ReturnResponse(
			c,
			http.StatusUnprocessableEntity,
			validate,
			nil,
			"validation error",
			"validation error")
	}
	fromDateParam, err := helper.ToStringDate(c.FormValue("from_date"))
	if err != nil {
		return helper.ReturnResponse(
			c,
			http.StatusUnprocessableEntity,
			validate,
			nil,
			"error convert",
			"error convert")
	}
	toDateParam, err := helper.ToStringDate(c.FormValue("to_date"))
	if err != nil {
		return helper.ReturnResponse(
			c,
			http.StatusUnprocessableEntity,
			validate,
			nil,
			"error convert",
			"error convert")
	}
	interval := helper.ToInt(c.FormValue("interval")) * 4
	userId := helper.ToInt(c.FormValue("user_id"))

	fromDate := helper.StartDay(fromDateParam).AddDate(0, 0, -interval)
	toDate := helper.EndDay(toDateParam)

	fromDateActual := helper.StartDay(fromDateParam).AddDate(0, 0, 0)

	response, responseActual, err := b.userService.GetData(c.Request().Context(), userId, interval,
		fromDate.Format("2006-01-02 15:04"), toDate.Format("2006-01-02 15:04"),
		fromDateActual.Format("2006-01-02 15:04"))
	if err != nil {
		return helper.ReturnInvalidResponse(http.StatusBadRequest, nil, nil, err.Error(), err.Error())
	}

	return helper.ReturnResponseData(c, http.StatusOK, response, responseActual, nil,
		"Success", "Success")

}

func (b *UserController) DataDefault(c echo.Context) error {

	rules := govalidator.MapData{
		"user_id": []string{"required", "numeric"},
	}

	validate := helper.ValidateRequestFormData(c, rules)
	if validate != nil {
		return helper.ReturnResponse(
			c,
			http.StatusUnprocessableEntity,
			validate,
			nil,
			"validation error",
			"validation error")
	}
	interval := 7 * 4
	userId := helper.ToInt(c.FormValue("user_id"))

	fromDate := helper.StartDay(time.Now()).AddDate(0, 0, -interval)
	toDate := helper.EndDay(time.Now())

	fromDateActual := helper.StartDay(time.Now()).AddDate(0, 0, -7)

	response, responseActual, err := b.userService.GetData(c.Request().Context(), userId, interval,
		fromDate.Format("2006-01-02 15:04"), toDate.Format("2006-01-02 15:04"),
		fromDateActual.Format("2006-01-02 15:04"))
	if err != nil {
		return helper.ReturnInvalidResponse(http.StatusBadRequest, nil, nil, err.Error(), err.Error())
	}

	return helper.ReturnResponseData(c, http.StatusOK, response, responseActual, nil,
		"Success", "Success")

}

func (b *UserController) Health(c echo.Context) error {

	return c.JSON(200, "Service is Running")

}
