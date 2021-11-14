package dto

import "time"

type LoginResponse struct {
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type SalesResponse struct {
	ID        int       `json:"id"`
	SalesDate time.Time `json:"sales_date"`
	Value     int       `json:"value"`
	UserId    int       `json:"user_id"`
}

type PayloadForecast struct {
	UserId int    `json:"user_id"`
	Period int    `json:"period"`
	Data   []Data `json:"data"`
}

type Data struct {
	DS string `json:"ds"`
	Y  int    `json:"y"`
}

type ResponseForecast struct {
	Messages []struct {
		DS     int     `json:"ds"`
		UserId int     `json:"user_id"`
		Yhat   float64 `json:"yhat"`
	} `json:"messages"`
}
