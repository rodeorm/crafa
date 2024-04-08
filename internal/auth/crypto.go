package auth

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// const MySecret string = "top secret key" //на проде заменить на генерацию случайной строки

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

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

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

// ReturnShortKey возвращает рандомный ключ (используем для генерации URL, подтверждающих сочетание адреса электронной почты и логина пользователя)
func ReturnShortKey(n int) (string, error) {
	if n <= 0 {
		err := fmt.Errorf("некорректное значение ключа %v", n)
		return "", err
	}
	rand.New(rand.NewSource((time.Now().UnixNano())))
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b), nil
}
