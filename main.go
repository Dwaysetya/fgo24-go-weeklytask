package main

import (
	"weeklytask/auth"
	"weeklytask/menu"
	"weeklytask/models"
)

func main() {
	service := auth.NewAuthHandler(&models.Users)
	auth.PromptAuth(service)
	menu.TampilkanMenu()
}
