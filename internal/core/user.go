package core

import "context"

// User - это сесияя для пользователя
type User struct {
	ID         int
	Login      string // Логин
	Surname    string // Фамилия
	Name       string // Имя
	Patronymic string // Отчество
	Email      string // Адрес электронной почты
	Phone      string // Телефон
	Password   string // Пароль
	Role       Role

	RoleID int
}

type UserStorager interface {
	RegUser(ctx context.Context, u *User) (*Session, error)               //RegUser добавляет нового пользователя. Возвращает письмо для подтверждения адреса электронной почты и сессию
	ConfirmUserEmail(ctx context.Context, login, email, otp string) error //ConfirmUserEmail подтверждает адрес электронной почты для нового пользователя. Возвращает письмо для подтверждения адреса электронной почты и сессию
	BaseAuthUser(context.Context, *User) error                            //BaseAuthUser авторизует пользователя через базовую аутентификацию по логину-паролю
	AdvAuthUser(context.Context, *User, string) error                     //AdvAuthUser авторизует пользователя, прошедшего базовую аутентификацию по одноразовому паролю
	UpdateUser(context.Context, *User) error                              //UpdateUser обновляет данные пользователя
}
