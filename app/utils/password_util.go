package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	cost := 12
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPasswordHash(password string) bool {
	hashedPassword, err := HashPassword(password)

	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func CheckPasswordAndConfirmPassword(password string, ConfirmPassword string) bool {
	return password == ConfirmPassword
}
