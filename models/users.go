package models

type User struct {
	Nama     string
	Email    string
	Password string
}

// Method untuk mendapatkan nama user
func (u User) GetNama() string {
	return u.Nama
}

// Method untuk mendapatkan email user
func (u User) GetEmail() string {
	return u.Email
}

// Method untuk memverifikasi password
func (u User) CheckPassword(password string) bool {
	return u.Password == password
}

var Users []User
