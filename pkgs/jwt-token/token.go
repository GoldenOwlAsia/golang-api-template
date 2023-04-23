package jwt_token

import (
	"fmt"
	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// GenerateAccessToken  generates a new JWT access token using the given userId.
func GenerateAccessToken(userId string, key string) (string, error) {
	// define token claims
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(configs.AccessTokenExpireTime).Unix(), // set token to expire in 15 mins
	}

	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token with secret key
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	// return signed token
	return signedToken, nil
}

func GenerateRefreshToken(userId, key string) (string, error) {
	// define token claims
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(configs.RefreshTokenExpireTime).Unix(), // set token to expire in 24 hours
	}

	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token with secret key
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	// return signed token
	return signedToken, nil
}

func RefreshAccessToken(refreshTokenString string) (string, error) {
	// parse refresh token
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return secret key
		return []byte(configs.ConfApp.SecretKey), nil
	})

	// check for errors
	if err != nil {
		return "", err
	}

	// get user ID from token
	userID, ok := refreshToken.Claims.(jwt.MapClaims)["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}

	// generate new access token
	accessToken, err := GenerateAccessToken(userID, configs.ConfApp.SecretKey)
	if err != nil {
		return "", err
	}

	// return access token
	return accessToken, nil
}

func ValidateAccessToken(tokenString string, key string) (*jwt.Token, error) {
	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return secret key
		return []byte(key), nil
	})

	// check for errors
	if err != nil {
		return nil, err
	}

	// check token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// return token
	return token, nil
}
