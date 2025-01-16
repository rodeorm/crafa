package core

import "context"

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
}

type UserStorager interface {
	RegUser(context.Context, *User) error
	AuthUser(context.Context, *User) error
	UpdateUser(context.Context, *User) error
}
