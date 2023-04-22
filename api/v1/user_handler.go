package v1

import (
	"api/api/v1/requests"
	"api/configs"
	"api/pkgs/jwt_auth_token"
	"api/services"
	"api/utils"
	"net/http"
	"os"

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

// Create		 godoc
// @Summary      Create user
// @Description  Create by username, password, email
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        json body requests.UserCreateRequest true "body params"
// @Success      201  {object}  utils.ResponseSuccess{Data=gorms.User}
// @Failure      422  {object}  utils.ResponseFailed{}
// @Failure      500  {object}  utils.ResponseFailed{}
// @Router       /api/v1/user [post]
func (receiver UserHandler) Create(c *gin.Context) {
	var req requests.UserCreateRequest
	errBind := c.Bind(&req)
	if errBind != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.GetRespError("invalid params", nil))
		return
	}

	res, err := receiver.service.Create(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetRespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, utils.GetRespSuccess("ok", res))
}

// GetByUsername godoc
// @Summary      Get user by username
// @Description  GetByUsername by username
// @Tags         users
// @Produce      json
// @Param		 Authorization header string true "Authorization"
// @Param        username query string true "username param"
// @Success      200  {object}  utils.ResponseSuccess{Data=gorms.User}
// @Failure      500  {object}  utils.ResponseFailed{}
// @Router       /api/v1/user [get]
func (receiver UserHandler) GetByUsername(c *gin.Context) {
	username := c.Query("username")

	res, err := receiver.service.GetByUsername(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetRespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.GetRespSuccess("ok", res))
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
func (receiver UserHandler) Login(c *gin.Context) {
	var req requests.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.GetRespError("invalid input", nil))
		return
	}

	resLogin, err := receiver.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.GetRespError(err.Error(), nil))
		return
	}

	//c.SetCookie("token", resLogin.AccessToken, configs.ConfApp.TokenMaxAge*60, "/", configs.ConfApp.Domain, false, true)

	c.JSON(http.StatusOK, utils.GetRespSuccess("welcome back", resLogin))
}

// Logout godoc
// @Summary      Logout user
// @Description
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Authorization"
// @Success      200  {object}  utils.ResponseSuccess{}
// @Failure      400  {object}  utils.ResponseFailed{}
// @Router       /api/v1/user/logout [post]
func (receiver UserHandler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", configs.ConfApp.Domain, false, true)

	c.JSON(http.StatusOK, utils.GetRespSuccess("ok", nil))
}

func (receiver UserHandler) GenerateTokenHandler(c *gin.Context) {
	// get user ID from context
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found"})
		return
	}

	// generate token
	token, err := jwt_auth_token.GenerateAccessToken(userId.(string), os.Getenv("SECRET_KEY"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// return token
	c.JSON(http.StatusOK, token)
}

func (receiver UserHandler) RefreshAccessTokenHandler(c *gin.Context) {
	// get refresh token from request
	refreshTokenString, ok := c.GetPostForm("refresh_token")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token not found"})
		return
	}

	// refresh access token
	accessToken, err := jwt_auth_token.RefreshAccessToken(refreshTokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return access token
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
