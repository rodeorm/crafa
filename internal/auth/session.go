package auth

import (
	"fmt"
	"money/internal/core"
)

func (s postgresStorage) SelectActiveSession(userID, sessionID int) (*core.Session, error) {
	if userID == 0 || sessionID == 0 {
		return nil, fmt.Errorf("некорректный пользователь или сессия")
	}
	session := core.Session{User: &core.User{ID: -1}}
	stmt, _ := s.DB.Preparex("SELECT ID, UseriD, D_Login, D_Last FROM cmn.Sessions WHERE ID = $1 AND D_Logout IS NULL")
	stmt.QueryRow(sessionID).Scan(&session.ID, &session.User.ID, &session.LogInTime, &session.LastActionTime)

	if session.ID == 0 {
		return nil, fmt.Errorf("нет сессии с таким идентификатором %d", sessionID)
	}

	if userID != session.User.ID {
		return nil, fmt.Errorf("сессия с идентификатор %d не принадлежит пользователю с идентификатором %d", sessionID, userID)
	}

	return &session, nil
}

func (s postgresStorage) SelectSession(userID, sessionID int) (*core.Session, error) {
	stmt, _ := s.DB.Preparex("SELECT ID, UserID, D_Login, D_Last, D_Logout FROM cmn.Sessions WHERE ID = $1")
	session := core.Session{User: &core.User{ID: 0}}
	stmt.QueryRow(sessionID).Scan(&session.ID, &session.User.ID, &session.LogInTime, &session.LastActionTime, &session.LogOutTime)

	if session.ID == 0 {
		return nil, fmt.Errorf("нет сессии с таким идентификатором")
	}
	if userID != session.User.ID {
		return nil, fmt.Errorf("сессия с идентификатор %d не принадлежит пользователю с идентификатором %d", sessionID, userID)
	}
	return &session, nil
}

func (s postgresStorage) UpdateSession(*core.Session) error {
	return nil
}

func (s postgresStorage) AddSession(user *core.User) (int, error) {
	stmt, _ := s.DB.Preparex("INSERT INTO cmn.Sessions (UserID, D_Login, D_Last) " +
		"SELECT $1, NOW(), NOW() RETURNING ID")

	var lastInsertID int
	err := stmt.QueryRow(user.ID).Scan(&lastInsertID)
	if lastInsertID == 0 {
		fmt.Println("AddSession", err)
		return 0, err
	}

	return lastInsertID, nil
}
