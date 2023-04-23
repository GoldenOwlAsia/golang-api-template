package handlers

import (
	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"github.com/GoldenOwlAsia/golang-api-template/handlers/requests"
	jwt_token "github.com/GoldenOwlAsia/golang-api-template/pkgs/jwt-token"
	"github.com/GoldenOwlAsia/golang-api-template/services"
	"github.com/GoldenOwlAsia/golang-api-template/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h UserHandler) GenerateTokenHandler(c *gin.Context) {
	// get user ID from context
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found"})
		return
	}

	// generate token
	token, err := jwt_token.GenerateAccessToken(userID.(string), configs.ConfApp.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// return token
	c.JSON(http.StatusOK, token)
}

func (h *UserHandler) RefreshAccessTokenHandler(c *gin.Context) {
	// get refresh token from request
	refreshTokenString, ok := c.GetPostForm("refresh_token")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token not found"})
		return
	}

	// refresh access token
	accessToken, err := jwt_token.RefreshAccessToken(refreshTokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return access token
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
