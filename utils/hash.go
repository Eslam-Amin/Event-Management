package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPasswordBytes), err
}