package core

import (
	"context"
	"database/sql"
	"time"
)

type Session struct {
	ID int
	User

	LoginTime      time.Time
	LogoutTime     sql.NullTime
	LastActionTime sql.NullTime
}

type SessionStorager interface {
	AddSession(context.Context, *User) (*Session, error) // На основании данных пользователя добавляем сесиию
	UpdateSession(context.Context, *Session) error
}
