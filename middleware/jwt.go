package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"github.com/GoldenOwlAsia/golang-api-template/models"
	jwtToken "github.com/GoldenOwlAsia/golang-api-template/pkgs/jwt-token"
	"github.com/GoldenOwlAsia/golang-api-template/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type jwtMiddleware struct {
	db *gorm.DB
}

func NewJwtAuth(db *gorm.DB) *jwtMiddleware {
	return &jwtMiddleware{db: db}
}

func (m jwtMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		token, err := extractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.GetRespError(err.Error(), nil))
			return
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.GetRespError("you are not logged in", nil))
			return
		}

		sub, err := jwtToken.ValidateAccessToken(token, configs.ConfApp.SecretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.GetRespError("your token is invalid", nil))
			return
		}

		userID := cast.ToUint(fmt.Sprint(sub))

		var user models.User
		err = m.db.Where(&models.User{ID: userID}).First(&user).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetRespError("the user belonging to this token no logger exists", nil))
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	token := strings.Split(header, " ")
	if len(token) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return token[1], nil
}
