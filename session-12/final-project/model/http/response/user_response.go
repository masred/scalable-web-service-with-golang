package response

import "time"

type UserRegisterResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateResponse struct {
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDeleteResponse struct {
	Message string `json:"message"`
}
