package utils

import "golang.org/x/crypto/bcrypt"

func HasPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 1) //should be 14
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
