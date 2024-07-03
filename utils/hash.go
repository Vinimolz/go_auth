package utils

import "golang.org/x/crypto/bcrypt"

func CreateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(userPassword, existingUserHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(existingUserHash), []byte(userPassword))

	return err == nil
}