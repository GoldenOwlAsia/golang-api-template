package middleware

import (
	"fmt"
	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"github.com/GoldenOwlAsia/golang-api-template/models"
	jwt_token "github.com/GoldenOwlAsia/golang-api-template/pkgs/jwt-token"
	"github.com/GoldenOwlAsia/golang-api-template/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type jwtMiddleware struct {
	db *gorm.DB
}

func NewJwtMiddleware(db *gorm.DB) *jwtMiddleware {
	return &jwtMiddleware{db: db}
}

func (m jwtMiddleware) DeserializeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		cookie, err := c.Cookie("token")

		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.GetRespError("you are not logged in", nil))
			return
		}

		sub, err := jwt_token.ValidateAccessToken(token, configs.ConfApp.SecretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.GetRespError("your token is invalid", nil))
			return
		}

		userId, _ := strconv.Atoi(fmt.Sprint(sub))

		var user models.User
		err = m.db.Where(&models.User{ID: uint(userId)}).First(&user).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetRespError("the user belonging to this token no logger exists", nil))
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}
