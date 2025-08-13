package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/SamJohn04/notes-backend/internal/config"
)

const userIdKey = "userId"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id int

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.JWTSecret), nil
		})
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusBadGateway)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			http.Error(w, "Something went wrong", http.StatusBadGateway)
			return
		}

		err = config.DB.QueryRow("SELECT id FROM users WHERE email=?", email).Scan(&id)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIdKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserId(r *http.Request) (int, error) {
	id, ok := r.Context().Value(userIdKey).(int)
	if !ok {
		return -1, errors.New("Not a valid id")
	}
	return id, nil
}
