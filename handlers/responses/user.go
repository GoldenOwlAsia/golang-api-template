package responses

import "github.com/GoldenOwlAsia/golang-api-template/models"

type UserLoginResponse struct {
	User         models.User `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
