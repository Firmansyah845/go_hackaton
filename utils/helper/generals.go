package helper

import (
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	TotalPage   int `json:"total_page"`
}

func ReturnInvalidResponse(httpcode int, data, meta interface{}, messageClient, messageServer string) error {
	responseBody := map[string]interface{}{
		"data": data,
		"meta": meta,
		"status": map[string]interface{}{
			"code":           httpcode,
			"message_client": messageClient,
			"message_server": messageServer,
		},
	}

	return echo.NewHTTPError(httpcode, responseBody)
}

func ReturnResponse(c echo.Context, httpcode int, data, meta interface{}, messageClient, messageServer string) error {
	responseBody := map[string]interface{}{
		"data": data,
		"meta": meta,
		"status": map[string]interface{}{
			"code":           httpcode,
			"message_client": messageClient,
			"message_server": messageServer,
		},
	}

	return c.JSON(httpcode, responseBody)
}

func ValidateRequestFormData(c echo.Context, rules govalidator.MapData) (i interface{}) {
	opts := govalidator.Options{
		Request: c.Request(),
		Rules:   rules,
	}

	v := govalidator.New(opts)
	mappedError := v.Validate()
	if len(mappedError) > 0 {
		i = mappedError
	}

	return i
}
