package middleware

import (
	"api/configs"
	"api/models/gorms"
	"api/pkgs/jwt"
	"api/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

		sub, errCouria := jwt.ValidateToken(token, configs.ConfApp.TokenSecret)
		if errCouria.Error != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.GetRespError("your token is invalid", nil))
			return
		}

		userId, _ := strconv.Atoi(fmt.Sprint(sub))

		var user gorms.User
		result := receiver.db.Where(&gorms.User{Id: uint(userId)}).First(&user)
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
