package main

import (
	"fmt"
)

func generatePassword(length int) string {
	return ""
}

func main() {
	var length int
	fmt.Print("\033[35;1mEnter the PASSWORD LENGHT: \033[0m")
	fmt.Scan(&length)

	password := generatePassword(length)
	fmt.Println("\033[32;1mGenerated PASSWORD:\033[0m", password)
}
