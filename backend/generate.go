package backend

import (
	"math/rand"
)

func GenerateRandomPassword(length int) string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums := "1234567890"
	chars := "!@#$%^&"
	allChars := upper + lower + nums + chars

	// rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for c := range password {
		password[c] = allChars[rand.Intn(len(allChars))]
	}
	return string(password)
}
