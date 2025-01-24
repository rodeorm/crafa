package postgres

import (
	"context"
	"log"
	"money/internal/core"
	"money/internal/crypt"
	"time"
)

// RegUser создает пользователя в БД
func (s *postgresStorage) RegUser(ctx context.Context, u *core.User, domain string) (*core.Session, error) {
	passwordHash, err := crypt.HashPassword(u.Password)
	if err != nil {
		log.Println("RegUser 1", err)
		return nil, err
	}

	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("RegUser 2", err)
		return nil, err
	}

	err = tx.Stmtx(s.preparedStatements["insertUser"]).GetContext(ctx, &u.ID, core.RoleReg, u.Login, passwordHash, u.Name, u.FamilyName, u.PatronName, u.Email, u.Phone)
	if err != nil {
		log.Println("RegUser 3", err)
		tx.Rollback()
		return nil, err
	}

	m := core.NewConfirmEmail(*u, domain)
	err = tx.Stmtx(s.preparedStatements["insertEmail"]).GetContext(ctx, &m.ID, core.MessageConfirm, u.ID, m.Text, m.Email)
	if err != nil {
		log.Println("RegUser 4", err)
		tx.Rollback()
		return nil, err
	}

	u.Role = core.Role{ID: core.RoleReg}

	session := &core.Session{User: *u}
	err = tx.Stmtx(s.preparedStatements["insertSession"]).GetContext(ctx, &session.ID, u.ID, time.Now(), time.Now())
	if err != nil {
		log.Println("RegUser 5", err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return session, nil
}

// Аутентифицирует пользователя на основании данных в БД и возвращает все его данные
func (s *postgresStorage) BaseAuthUser(ctx context.Context, u *core.User) bool {
	return false
}

// Аутентифицирует пользователя на основании данных в БД и возвращает все его данные
func (s *postgresStorage) GetUser(ctx context.Context, u *core.User) error {
	return s.preparedStatements["selectUser"].GetContext(ctx, u, u.ID)
}
