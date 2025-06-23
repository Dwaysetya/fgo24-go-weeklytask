package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptAuth(service AuthService) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(`
=================================
      🔐 AUTHENTICATION MENU      
=================================
1. 🔑 Login
2. 📝 Register
0. ❌ Keluar
=================================
`)
		fmt.Print("Pilih opsi: ")
		var opsi int
		fmt.Scanln(&opsi)

		switch opsi {
		case 1:
			fmt.Print("Email: ")
			email, _ := reader.ReadString('\n')
			fmt.Print("Password: ")
			password, _ := reader.ReadString('\n')

			user := service.Login(strings.TrimSpace(email), strings.TrimSpace(password))
			if user != nil {
				fmt.Println("✅ Login berhasil sebagai", user.Nama)
				return
			}
			fmt.Println("❌ Email atau password salah")
		case 2:
			fmt.Print("Nama: ")
			nama, _ := reader.ReadString('\n')
			fmt.Print("Email: ")
			email, _ := reader.ReadString('\n')
			fmt.Print("Password: ")
			password, _ := reader.ReadString('\n')

			service.Register(
				strings.TrimSpace(nama),
				strings.TrimSpace(email),
				strings.TrimSpace(password),
			)
			fmt.Println("✅ Registrasi berhasil!")
		case 0:
			fmt.Println("👋 Keluar dari program.")
			os.Exit(0)
		default:
			fmt.Println("❌ Pilihan tidak valid")
		}
	}
}
