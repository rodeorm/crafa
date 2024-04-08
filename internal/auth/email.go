package auth

import (
	"context"
	"fmt"
	"time"

	"money/internal/core"
)

func (s *postgresStorage) SelectEmailsForSending(ctx context.Context) (*[]core.Email, error) {
	emails := make([]core.Email, 0, 1)
	err := s.DB.SelectContext(ctx, &emails, "SELECT C_key AS key, C_email AS email, C_login AS login FROM cmn.emails WHERE D_Sended IS NULL")
	if err != nil {
		return nil, err
	}
	return &emails, nil
}

func (s *postgresStorage) UpdateEmail(ctx context.Context, em *core.Email) error {

	query := "UPDATE cmn.emails" +
		" SET C_key = $2, D_Sended = $3, D_Confirmed = $4" +
		" WHERE C_Email = $1"
	_, err := s.DB.ExecContext(ctx, query, em.Email, em.Key, em.SendedTime, em.ConfirmedTime)
	if err != nil {
		return err
	}
	return nil
}

func (s *postgresStorage) ApproveUserEmail(ctx context.Context, login string, key string) (int, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("Ошибка при объявлении транзакции:", err)
		return 0, err
	}
	var (
		test_login string
		userID     int
	)
	// Проверка комбинации логин + код для емэйла
	s.DB.QueryRowContext(ctx, "SELECT C_Login, userID from cmn.emails WHERE C_Login = $1 AND C_Key = $2", login, key).Scan(&test_login, &userID)
	if test_login == "" {
		return 0, fmt.Errorf("невалидная комбинация логина и ключа подтверждения")
	}
	// Обновление даты подтверждения емэйла
	query := "UPDATE cmn.emails" +
		" SET D_Confirmed = $3" +
		" WHERE C_login = $1 AND C_key = $2"
	_, err = tx.ExecContext(ctx, query, login, key, time.Now())
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	// Изменение роли пользователя
	query = "UPDATE cmn.users" +
		" SET roleID = 3" +
		" WHERE ID = $1"
	_, err = tx.ExecContext(ctx, query, userID)

	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}
	return userID, nil
}
