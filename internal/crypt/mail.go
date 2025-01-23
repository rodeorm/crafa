package crypt

import (
	"crypto/rand"
	"encoding/base32"
)

// GetOneTimePassword возвращает одноразовый пароль
func GetOneTimePassword() string {
	// Создаем байтовый массив необходимой длины
	length := 6
	bytes := make([]byte, length)

	// Генерируем криптостойкие случайные байты
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}

	// Кодируем байты в строку Base32
	otp := base32.StdEncoding.EncodeToString(bytes)
	return otp[:length]
}

func GetVerifyURL(url string) string {
	// Создаем байтовый массив необходимой длины
	length := 6
	bytes := make([]byte, length)

	// Генерируем криптостойкие случайные байты
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}

	// Кодируем байты в строку Base32
	otp := base32.StdEncoding.EncodeToString(bytes)

	return url + "/" + otp[:length]
}
