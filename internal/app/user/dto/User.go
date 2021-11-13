package dto

type LoginResponse struct {
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type SalesResponse struct {
	ID        int    `json:"id"`
	SalesDate string `json:"sales_date"`
	Value     int    `json:"value"`
	UserId    int    `json:"user_id"`
}
