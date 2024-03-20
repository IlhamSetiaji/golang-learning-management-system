package response

type UserLoginResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
	Expires   int64  `json:"expires"`
}

type UserMeResponse struct {
	ID              uint           `json:"id"`
	Name            string         `json:"name"`
	Username        string         `json:"username"`
	Email           string         `json:"email"`
	EmailVerifiedAt string         `json:"email_verified_at"`
	Roles           []RoleResponse `json:"roles"`
}
