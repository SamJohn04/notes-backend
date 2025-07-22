package repository

import (
	"errors"
	"github.com/SamJohn04/notes-backend/internal/model"
)

var users = map[string]model.User{} // TODO change to DB

func CreateUser(user model.User) error {
	if _, exists := users[user.Email]; exists {
		return errors.New("user already exists")
	}
	users[user.Email] = user
	return nil
}

func GetUserByEmail(email string) (model.User, error) {
	user, exists := users[email]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}
