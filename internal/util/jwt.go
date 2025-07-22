package util

import (
	"github.com/SamJohn04/notes-backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte(config.Cfg.JWTSecret)

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(jwtKey)
}
