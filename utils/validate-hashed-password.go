package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func ValidateHashedPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return fmt.Errorf("password is wrong %v", err)
	}

	return nil
}
