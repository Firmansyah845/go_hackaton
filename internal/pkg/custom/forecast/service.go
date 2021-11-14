package forecast

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Firmansyah845/go_hackaton/config"
	"github.com/Firmansyah845/go_hackaton/internal/app/user/dto"
	"net/http"
	"time"
)

type ServiceForecastingImpl struct {
}

func (s ServiceForecastingImpl) GetDataForecasting(ctx context.Context, request []dto.LoginResponse) error {
	marshalled, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", config.App.Config.URL_FORECAST_SERVICE, bytes.NewBuffer(marshalled))
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func NewServiceMonetize() ServiceForecasting {
	return &ServiceForecastingImpl{}
}
