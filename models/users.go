package models

type User struct {
	Nama     string
	Email    string
	Password string
}

func (u User) GetNama() string {
	return u.Nama
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) CheckPassword(password string) bool {
	return u.Password == password
}

var Users []User
