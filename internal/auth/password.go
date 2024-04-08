package auth

import (
	"context"
	"fmt"
	"log"

	"money/internal/core"
)

// AuthUser аутентифицирует Пользователя
func (s *postgresStorage) AuthUser(ctx context.Context, u *core.User) (success bool, err error) {

	success, err = s.CheckPassword(ctx, u.Login, u.Password)

	if success && err == nil {
		s.SelectUser(ctx, u)
		return success, err
	}
	if err != nil {
		log.Println("AuthUser", err)
		return false, err
	}
	return false, nil
}

// CheckPassword проверяет пароль
func (s postgresStorage) CheckPassword(ctx context.Context, login, password string) (bool, error) {
	var passwordDB string
	err := s.DB.QueryRowContext(ctx, "SELECT C_Password FROM cmn.Users WHERE C_Login = $1; ", login).Scan(&passwordDB)
	if err != nil {
		log.Println("CheckPassword", err)
		return false, fmt.Errorf("неправильно указан логин или пароль")
	}

	if CheckPasswordHash(password, passwordDB) {
		return true, nil
	}
	return false, nil
}

// Смена пароля пользователя
func (s *postgresStorage) ChangePassword(ctx context.Context, u *core.User) error {
	hashedPassword, _ := HashPassword(u.Password)
	query := "UPDATE cmn.users" +
		" SET C_password = $2" +
		" WHERE userID = $1"
	_, err := s.DB.ExecContext(ctx, query, u.ID, hashedPassword)
	if err != nil {
		log.Println("ChangePassword", err)
		return err
	}
	return nil
}
