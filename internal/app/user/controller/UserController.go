package controller

import (
	"github.com/Firmansyah845/go_hackaton/internal/app/user/service"
	"github.com/Firmansyah845/go_hackaton/utils/helper"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"net/http"
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

func (b *UserController) Health(c echo.Context) error {

	return c.JSON(200, "Service Rule Engine is Running")

}
