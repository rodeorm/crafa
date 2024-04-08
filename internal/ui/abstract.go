package ui

import (
	"context"
	"money/internal/core"
)

type AuthRepo interface {
	// Зарегистрировать нового пользователя
	RegUser(ctx context.Context, u *core.User) error
	// Аутентфиицировать пользователя
	AuthUser(context.Context, *core.User) (bool, error)
	// Обновить данные пользователя
	UpdateUser(ctx context.Context, u *core.User) error
	// Получить данные пользователя
	SelectUser(ctx context.Context, u *core.User) error
	// Удалить/Деактивировать пользователя
	DeleteUser(ctx context.Context, u *core.User) error

	// Выбрать активную сессию
	SelectActiveSession(int, int) (*core.Session, error)
	// Обновить данные сессии
	UpdateSession(*core.Session) error
	// Добавить новую сессию
	AddSession(*core.User) (int, error)
	// Получить данные сессии
	SelectSession(int, int) (*core.Session, error)

	// Разорвать соединение
	CloseConnection()
}

type Log interface {
}

type WorkPlaceRepo interface {
	CloseConnection()
}
