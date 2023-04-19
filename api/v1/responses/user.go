package responses

type UserLoginResponse struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}
