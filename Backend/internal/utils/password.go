package utils

import (
	"golang.org/x/crypto/bcrypt"
	"os"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("SALT")+password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(os.Getenv("SALT")+password))
	return err == nil
}
