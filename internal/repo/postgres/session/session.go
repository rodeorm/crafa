package session

import (
	"context"
	"crafa/internal/core"
)

// StartSession начинает новую сессию
func StartSession(context.Context, *core.User) (*core.Session, error) {
	return nil, nil
}

// UpdateSession обновляет данные сессии
func UpdateSession(context.Context, *core.Session) error {
	return nil
}

// EndSession закрывает сессию
func EndSession(context.Context, *core.Session) error {
	return nil
}
