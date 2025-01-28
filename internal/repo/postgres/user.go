package postgres

import (
	"context"
	"fmt"
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

	_, err = tx.Stmtx(s.preparedStatements["insertMsg"]).ExecContext(ctx, core.MessageTypeConfirm, core.MessageCategoryEmail, u.ID, crypt.GetOneTimePassword(), u.Email)
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

func (s *postgresStorage) ConfirmUserEmail(ctx context.Context, userID int, otp string) error {
	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("ConfirmUserEmail 1", err)
		return err
	}
	msg := core.Message{}
	// Проверяем переданный код  UserID = $1 AND Text = $2 AND Email = $3
	err = tx.Stmtx(s.preparedStatements["selectConfMsg"]).GetContext(ctx, &msg, userID, otp)
	if err != nil {
		log.Println("ConfirmUserEmail 2", err)
		tx.Rollback()
		return err
	}

	if msg.ID == 0 {
		tx.Rollback()
		return fmt.Errorf("переданные данные невалидны. нет сообщения для такого пользователя с таким адресом и кодом подтверждения")
	}
	msg.Used = true

	// Обновляем сообщение, что оно было использовано
	_, err = tx.Stmtx(s.preparedStatements["updateMsg"]).ExecContext(ctx, msg.ID, msg.Used, msg.Queued, msg.SendTime)
	if err != nil {
		log.Println("ConfirmUserEmail 2", err)
		tx.Rollback()
		return err
	}

	// Обновляем роль пользователя
	_, err = tx.Stmtx(s.preparedStatements["updateUserRole"]).ExecContext(ctx, userID, core.RoleAuth)
	if err != nil {
		log.Println("ConfirmUserEmail 3", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// Аутентифицирует пользователя на основании данных в БД и возвращает все его данные
func (s *postgresStorage) BaseAuthUser(ctx context.Context, u *core.User) bool {
	return false
}

// Аутентифицирует пользователя на основании данных в БД и возвращает все его данные
func (s *postgresStorage) GetUser(ctx context.Context, u *core.User) error {
	return s.preparedStatements["selectUser"].GetContext(ctx, u, u.ID)
}
