package auth

import (
	"sync"
	"weeklytask/models"
)

func (a *AuthHandle) Register(nama, email, password string) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		user := models.User{Nama: nama, Email: email, Password: password}
		*a.Users = append(*a.Users, user)
	}()

	wg.Wait()
}
