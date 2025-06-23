package auth

import (
	"sync"
	"weeklytask/models"
)

type AuthHandler struct {
	Users *[]models.User
}

func NewAuthHandle(users *[]models.User) *AuthHandler {
	return &AuthHandler{Users: users}
}

func (a *AuthHandler) Register(nama, email, password string) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		user := models.User{Nama: nama, Email: email, Password: password}
		*a.Users = append(*a.Users, user)
	}()

	wg.Wait()
}

func (a *AuthHandler) Login(email, password string) *models.User {
	result := make(chan *models.User)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for _, u := range *a.Users {
			if u.Email == email && u.Password == password {
				result <- &u
				return
			}
		}
		result <- nil
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	return <-result
}
