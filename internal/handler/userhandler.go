package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SamJohn04/notes-backend/internal/model"
	"github.com/SamJohn04/notes-backend/internal/repository"
	"github.com/SamJohn04/notes-backend/internal/util"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	hashed, err := util.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusBadGateway)
		return
	}
	user := model.User{Email: req.Email, Password: hashed}

	err = repository.CreateUser(user)
	if err != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}
	if !util.CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := util.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusBadGateway)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
