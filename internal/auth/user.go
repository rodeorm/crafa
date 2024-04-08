package auth

import (
	"context"
	"fmt"
	"log"

	"money/internal/core"
)

func (s *postgresStorage) SelectAllUsers(ctx context.Context) (*[]core.User, error) {

	rows, err := s.DB.QueryContext(ctx, "SELECT u.ID, u.C_Login, u.C_Surname, u.C_Name, u.C_Patronymic, u.C_Email, u.C_Phone, u.RoleID FROM cmn.users AS u")
	if err != nil {
		log.Println("SelectAllUsers", err)
		return nil, err
	}
	if rows.Err() != nil {
		log.Println("SelectAllUsers", err)
		return nil, err
	}
	defer rows.Close()
	users := make([]core.User, 0, 1)
	for rows.Next() {
		var (
			user core.User
			role core.Role
		)
		err = rows.Scan(&user.ID, &user.Login, &user.Surname, &user.Name, &user.Patronymic, &user.Email, &user.Phone, &role.ID)
		if err != nil {
			log.Println("SelectAllUsers", err)
			return nil, err
		}
		err := s.SelectRoleData(ctx, &role)
		if err != nil {
			log.Println("SelectAllUsers", err)
			return nil, err
		}
		user.Role = role

		users = append(users, user)
	}

	return &users, nil
}

func (s *postgresStorage) SelectUser(ctx context.Context, u *core.User) error {

	query := "SELECT ID, C_Login, C_Surname, C_Name, C_Patronymic, C_Email, C_Phone, Roleid" +
		" FROM cmn.users" +
		" WHERE C_Login = $1 OR ID = $2"
	var role core.Role

	err := s.DB.QueryRowContext(ctx, query, u.Login, u.ID).Scan(&u.ID, &u.Login, &u.Surname, &u.Name,
		&u.Patronymic, &u.Email, &u.Phone, &role.ID)
	if err != nil {
		log.Println("SelectUser", err)
		return err
	}

	err = s.SelectRoleData(ctx, &role)
	if err != nil {
		log.Println("SelectUser", err)
		return err
	}
	u.Role = role
	return nil
}

// Обновление данных пользователя (кроме пароля)
func (s *postgresStorage) UpdateUser(ctx context.Context, u *core.User) error {
	// Сначала получаем текущие данные пользователя в БД для сравнения
	currentUser := core.User{Login: u.Login, ID: u.ID}
	s.SelectUser(ctx, &currentUser)
	// Проверяем, что пользователя с таким email нет в БД
	var mail string

	query := "UPDATE cmn.users" +
		" SET C_Surname = $2, C_name = $3, C_Patronymic = $4, C_Email = $5, C_Phone = $6, Roleid = $7" +
		" WHERE C_Login = $1"

	if currentUser.Email != u.Email { // В этом случае делаем пользователя неавторизованном и заставляем подтвердить email
		u.ID = currentUser.ID
		u.Login = currentUser.Login
		u.Role.ID = 2
		s.DB.QueryRowContext(ctx, "SELECT C_Login from cmn.users WHERE C_Email = $1 AND UserID != $2", u.Email, u.ID).Scan(&mail)
		if mail != "" {
			return fmt.Errorf("пользователь с адресом электронной почты %s уже зарегистрирован", u.Email)
		}
		// Проверяем, что email корректный
		if !IsEmailValid(u.Email) {
			return fmt.Errorf("некорректный адрес электронной почты %s", u.Email)
		}

		tx, err := s.DB.BeginTx(ctx, nil)
		if err != nil {
			log.Println("UpdateUser. Ошибка при объявлении транзакции:", err)
			return err
		}

		tx.ExecContext(ctx, query, u.Login, u.Surname, u.Name, u.Patronymic, u.Email, u.Phone, u.Role.ID)
		keyForConfirmation, err := ReturnShortKey(10)
		if err != nil {
			log.Println("UpdateUser. Ошибка при генерации случайного ключа для подтверждения: ", err)
			tx.Rollback()
			return err
		}

		_, err = tx.ExecContext(ctx, "INSERT INTO cmn.emails (C_Key, C_Login, C_Email, UserID) "+
			"SELECT $1, $2, $3, $4", keyForConfirmation, u.Login, u.Email, u.ID)
		if err != nil {
			log.Println("UpdateUser. Ошибка при вставке подтверждающего сообщения: ", err)
			tx.Rollback()
			return err
		}

		return tx.Commit()
	}

	_, err := s.DB.ExecContext(ctx, query, u.Login, u.Surname, u.Name, u.Patronymic, u.Email, u.Phone, u.Role.ID)
	if err != nil {
		return err
	}
	return s.SelectUser(ctx, u)
}

func (s *postgresStorage) DeleteUser(ctx context.Context, u *core.User) error {
	// В транзакции удаляются данные из всех, связанных таблиц. TODO: по мере расширения состава таблиц - актуализировать
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println("DeleteUser 1", err)
		return err
	}
	s.SelectUser(ctx, u)
	// Удаляем сообщения для подтверждения email для этого логина
	query := "DELETE FROM cmn.emails" +
		" WHERE userID = $1"
	_, err = tx.ExecContext(ctx, query, u.ID)
	if err != nil {
		log.Println("DeleteUser 2", err)
		tx.Rollback()
		return err
	}

	// Удаляем сессии для этого логина
	query = "DELETE FROM cmn.sessions" +
		" WHERE userID = $1"
	_, err = tx.ExecContext(ctx, query, u.ID)
	if err != nil {
		log.Println("DeleteUser 3", err)
		tx.Rollback()
		return err
	}

	// Удаляем пользователей
	query = "DELETE FROM cmn.users" +
		" WHERE ID = $1"
	_, err = tx.ExecContext(ctx, query, u.ID)

	if err != nil {
		log.Println("DeleteUser 4", err)
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("DeleteUser 5", err)
		tx.Rollback()
	}
	return nil
}

// RegUser создает нового Пользователя
func (s *postgresStorage) RegUser(ctx context.Context, u *core.User) error {
	var login, mail string

	// Проверяем, что пользователя с таким логином нет в БД
	s.DB.QueryRowContext(ctx, "SELECT C_Login from cmn.users WHERE C_Login = $1", u.Login).Scan(&login)
	if login != "" {
		return fmt.Errorf("пользователь с логином %s уже зарегистрирован", u.Login)
	}

	// Проверяем, что пользователя с таким email нет в БД
	s.DB.QueryRowContext(ctx, "SELECT C_login from cmn.users WHERE C_Email = $1", u.Email).Scan(&mail)
	if mail != "" {
		return fmt.Errorf("пользователь с адресом электронной почты %s уже зарегистрирован", u.Email)
	}

	if !isPhoneValid(u.Phone) {
		return fmt.Errorf("номер телефона %s введен некорректно", u.Phone)
	}
	if !IsEmailValid(u.Email) {
		return fmt.Errorf("некорректный адрес электронной почты %s", u.Email)
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("ошибка при объявлении транзакции на вставку %s", err)
	}

	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return fmt.Errorf("ошибка при обработке пароля %s", err)
	}

	err = tx.QueryRowContext(ctx, "INSERT INTO cmn.users (C_Login, C_Password, C_Surname, C_Name, C_Patronymic, C_Email, C_Phone, RoleId) "+
		//	"SELECT $1, $2, $3, $4, $5, $6, $7, $8 RETURNING ID", u.Login, hashedPassword, u.Surname, u.Name, u.Patronymic, u.Email, u.Phone, 2).Scan(&u.ID)
		"SELECT $1, $2, $3, $4, $5, $6, $7, $8 RETURNING ID", u.Login, hashedPassword, u.Surname, u.Name, u.Patronymic, u.Email, u.Phone, 3).Scan(&u.ID)
	if err != nil {
		log.Println("Ошибка при вставке записи в таблицу пользователей :", err)
		tx.Rollback()
		ctx.Err()
		return err
	}
	/*
		keyForConfirmation, err := ReturnShortKey(10)
		if err != nil {
			log.Println("Ошибка при генерации случайного ключа для подтверждения: ", err)
			tx.Rollback()
			return err
		}

			_, err = tx.ExecContext(ctx, "INSERT INTO cmn.emails (C_Key, C_Login, C_Email, UserID) "+
				"SELECT $1, $2, $3, $4", keyForConfirmation, u.Login, u.Email, u.ID)
			if err != nil {
				log.Println("Ошибка при вставке подтверждающего сообщения: ", err)
				tx.Rollback()
				return err
			}
	*/
	err = tx.Commit()
	if err != nil {
		log.Println("RegUser", err)
		tx.Rollback()
	}

	return s.SelectUser(ctx, u)
}
