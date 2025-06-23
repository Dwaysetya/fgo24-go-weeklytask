package auth

import "weeklytask/models"

type AuthHandle struct {
	Users *[]models.User
}

func NewAuthHandler(users *[]models.User) *AuthHandle {
	return &AuthHandle{Users: users}
}
