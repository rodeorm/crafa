package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

// const MySecret string = "top secret key" //на проде заменить на генерацию случайной строки

// HashPassword хэширует пароль
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash проверят соответствие пароля и хэша
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
