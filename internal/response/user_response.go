package response

type UserLoginResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
