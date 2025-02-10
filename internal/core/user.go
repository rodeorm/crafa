package core

import (
	"context"
	"time"
)

// User - это сесияя для пользователя
type User struct {
	ID       int    `db:"user.id"`
	Login    string // Логин
	Password string // Пароль

	FamilyName string // Фамилия
	Name       string // Имя
	PatronName string // Отчество

	Email string // Адрес электронной почты
	Phone string // Телефон

	Role Role
}

type UserStorager interface {
	RegUser(ctx context.Context, u *User, domain string) (*Session, error)       //	RegUser добавляет нового пользователя. Возвращает письмо для подтверждения адреса электронной почты и сессию
	SelectUser(ctx context.Context, u *User) error                               //	SelectUser возвращает данные пользователя
	ConfirmUserEmail(ctx context.Context, userID int, otp string) error          //	ConfirmUserEmail подтверждает адрес электронной почты для нового пользователя. Возвращает ошибку, если подтверждение не удалось
	BaseAuthUser(context.Context, *User) error                                   //	BaseAuthUser авторизует пользователя через базовую аутентификацию по логину-паролю
	AdvAuthUser(context.Context, *User, string, time.Duration) (*Session, error) //	AdvAuthUser авторизует пользователя, прошедшего базовую аутентификацию по одноразовому паролю
	UpdateUser(context.Context, *User) error                                     //	UpdateUser обновляет данные пользователя
	SelectAllUsers(ctx context.Context) ([]User, error)                          // SelectAllUsers возвращает данные всех пользователей
}
