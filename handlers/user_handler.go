package handlers

import (
	"net/http"

	"github.com/GoldenOwlAsia/golang-api-template/handlers/requests"
	jwt_token "github.com/GoldenOwlAsia/golang-api-template/pkgs/jwt-token"
	"github.com/GoldenOwlAsia/golang-api-template/services"
	"github.com/GoldenOwlAsia/golang-api-template/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) UserHandler {
	return UserHandler{
		service: s,
	}
}

// Login godoc
// @Summary      Login user to system
// @Description  login by username & password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        json body requests.UserLoginRequest true "body params"
// @Success      200  {object}  utils.ResponseSuccess{Data=responses.UserLoginResponse}
// @Failure      401  {object}  utils.ResponseFailed{}
// @Failure      422  {object}  utils.ResponseFailed{}
// @Router       /api/v1/user/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req requests.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.GetRespError("invalid input", nil))
		return
	}

	resLogin, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.GetRespError(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.GetRespSuccess("welcome back", resLogin))
}

func (h *UserHandler) RefreshAccessToken(c *gin.Context) {
	refreshTokenString, ok := c.GetPostForm("refresh_token")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token not found"})
		return
	}
	newAccessToken, newRefreshToken, err := jwt_token.RefreshAccessToken(refreshTokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jwt_token.Token{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})
}
