package middleware

import (
	"api/configs"
	"api/models"
	"api/pkgs/jwt"
	"api/utils"
	"fmt"
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

func (receiver jwtMiddleware) DeserializeUser() gin.HandlerFunc {
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

		sub, err := jwt.ValidateToken(token, configs.ConfApp.TokenSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.GetRespError("your token is invalid", nil))
			return
		}

		userId, _ := strconv.Atoi(fmt.Sprint(sub))

		var user models.User
		result := receiver.db.Where(&models.User{Id: uint(userId)}).First(&user)
		if result.Error != nil {
			resp := utils.ResponseFailed{
				Code:       0,
				StatusCode: utils.ResponseStatusCodeFailed,
				Message:    "the user belonging to this token no logger exists",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}
