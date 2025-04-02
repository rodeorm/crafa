package postgres

import (
	"context"
	"money/internal/core"
)

// StartSession начинает новую сессию
func (s *PostgresStorage) StartSession(context.Context, *core.User) (*core.Session, error) {
	return nil, nil
}

// UpdateSession обновляет данные сессии
func (s *PostgresStorage) UpdateSession(context.Context, *core.Session) error {
	return nil
}

// EndSession закрывает сессию
func (s *PostgresStorage) EndSession(context.Context, *core.Session) error {
	return nil
}
