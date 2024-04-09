package request

type UserLoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type VerifyUserRequest struct {
	Token string `validate:"required"`
}

type RegisterUserRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	RoleID   uint   `json:"role_id" form:"role" binding:"required"`
}

type SendEmailRequest struct {
	Email string `json:"email" form:"email" binding:"required"`
}
