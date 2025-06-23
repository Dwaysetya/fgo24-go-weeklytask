package auth

import (
	"sync"
	"weeklytask/models"
)

func (a *AuthHandle) Login(email, password string) *models.User {
	resultChan := make(chan *models.User)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for _, u := range *a.Users {
			if u.Email == email && u.Password == password {
				resultChan <- &u
				return
			}
		}
		resultChan <- nil
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return <-resultChan
}
