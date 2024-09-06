package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pasword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
