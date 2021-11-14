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

func (s ServiceForecastingImpl) GetDataForecasting(ctx context.Context, request dto.PayloadForecast) (*dto.ResponseForecast, error) {
	var result dto.ResponseForecast

	marshalled, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", config.App.Config.URL_FORECAST_SERVICE, bytes.NewBuffer(marshalled))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewServiceMonetize() ServiceForecasting {
	return &ServiceForecastingImpl{}
}
