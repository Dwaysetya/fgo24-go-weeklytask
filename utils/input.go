package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		num, err := strconv.Atoi(strings.TrimSpace(input))
		if err == nil {
			return num
		}
		fmt.Println("Input tidak valid, Coba lagi")
	}
}
