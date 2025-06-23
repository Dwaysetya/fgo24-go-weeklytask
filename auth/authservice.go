package auth

import "weeklytask/models"

type AuthService interface {
	Register(nama, email, password string)
	Login(email, password string) *models.User
}
