package request

type UserRegisterRequest struct {
	Username string `binding:"required"  json:"username"`
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=6" json:"password"`
	Age      int    `binding:"required,gt=8" json:"age"`
}

type UserLoginRequest struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

type UserUpdateRequest struct {
	Email    string `binding:"required,email" json:"email"`
	Username string `binding:"required" json:"username"`
}
