package dto

type LoginResponse struct {
	UserId             int     `json:"user_id"`
	Name 			   string `json:"name"`
	Username           string `json:"username"`
	Role 			   string `json:"role"`
}
