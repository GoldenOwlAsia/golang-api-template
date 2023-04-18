package jwt

import (
	"api/models"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(user models.User, expirationTime time.Time, secretJWTKey string) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = user.Username
	claims["exp"] = expirationTime.Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err = token.SignedString([]byte(secretJWTKey))

	if err != nil {
		return "", errors.New("generate token failed")
	}

	return
}

func ValidateToken(token string, signedJWTKey string) (validateResult interface{}, err error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})

	if err != nil {
		return nil, errors.New("invalidate token")
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, errors.New("invalid token claim")
	}

	return claims["sub"], nil
}

func GetCurrentUser(c *gin.Context) (user models.User) {
	currentUser := c.MustGet("currentUser").(models.User)

	user = models.User{
		Username:       currentUser.Username,
		Password:       currentUser.Password,
		Email:          currentUser.Email,
		Role:           currentUser.Role,
		Status:         currentUser.Status,
		ApprovedStatus: currentUser.ApprovedStatus,
	}
	return
}
