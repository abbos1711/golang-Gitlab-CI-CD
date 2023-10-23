package utils

import (
	"math/rand"
	"time"
)

const (
	digitBytes = "0123456789"
)

func RandomPassword() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random password
	password := make([]byte, 6)
	for i := range password {
		password[i] = digitBytes[rand.Intn(len(digitBytes))]
	}

	return string(password)
}