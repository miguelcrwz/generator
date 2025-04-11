package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePassword(length int) string {
	if length < 2 {
		return "\033[31;1m✘ PASSWORD LENGHT MUST BE AT LEAST 2\033[0m"
	}

	rand.Seed(time.Now().UnixNano())

	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"

	allChars := lowerCase + upperCase + numbers + special

	mandatory := []byte{
		upperCase[rand.Intn(len(upperCase))],
		numbers[rand.Intn(len(numbers))],
	}

	password := make([]byte, length-len(mandatory))

	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	password = append(password, mandatory...)

	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

func main() {
	var length int
	fmt.Print("\033[35;1mEnter the PASSWORD LENGHT: \033[0m")
	fmt.Scan(&length)

	password := generatePassword(length)
	fmt.Println("\033[32;1mGenerated PASSWORD:\033[0m", password)
}
