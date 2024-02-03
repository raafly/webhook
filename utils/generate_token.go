package utils

import (
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raaafly/powerup-client-service-golang/utils/constans"
)

type GenerateToken struct{}

func NewGenerateToken() *GenerateToken {
	return &GenerateToken{}
}

func (g *GenerateToken) GenerateAccessToken(id, email, username string) (string, int64, error) {
	secret := os.Getenv("JWT_SECRET")
	expired := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"email":    email,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, constans.NewBadRequestError(err.Error())
	}

	return tokenString, expired, nil
}

func (g *GenerateToken) GenerateRefreshToken(id, email, username string) (string, int64, error) {
	secret := os.Getenv("JWT_SECRET")

	expired := time.Now().Add(time.Hour * 24 * 7).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"email":    email,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, constans.NewBadRequestError(err.Error())
	}

	return tokenString, expired, nil
}

func GenerateRandomCode(length int) string {
	rand.NewSource(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}