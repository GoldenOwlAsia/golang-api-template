package requests

type UserCreateRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
