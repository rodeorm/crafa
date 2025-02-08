package core

import (
	"context"
	"time"
)

// User - это сесияя для пользователя
type User struct {
	ID         int    `db:"user.id"`
	Login      string // Логин
	FamilyName string // Фамилия
	Name       string // Имя
	PatronName string // Отчество
	Email      string // Адрес электронной почты
	Phone      string // Телефон
	Password   string // Пароль
	Role       Role

	RoleID int
}

type UserStorager interface {
	RegUser(ctx context.Context, u *User, domain string) (*Session, error)       //RegUser добавляет нового пользователя. Возвращает письмо для подтверждения адреса электронной почты и сессию
	GetUser(ctx context.Context, u *User) error                                  //GetUser возвращает данные пользователя из БД
	ConfirmUserEmail(ctx context.Context, userID int, otp string) error          //ConfirmUserEmail подтверждает адрес электронной почты для нового пользователя. Возвращает ошибку, если подтверждение не удалось
	BaseAuthUser(context.Context, *User) error                                   //BaseAuthUser авторизует пользователя через базовую аутентификацию по логину-паролю
	AdvAuthUser(context.Context, *User, string, time.Duration) (*Session, error) //AdvAuthUser авторизует пользователя, прошедшего базовую аутентификацию по одноразовому паролю
	// UpdateUser(context.Context, *User) error                              //UpdateUser обновляет данные пользователя
}
