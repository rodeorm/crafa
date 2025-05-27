package core

import (
	"testing"
	"time"
)

func TestGetSessionFromTkn(t *testing.T) {
	jwtKey := "test-key"
	tokenLiveTime := time.Hour
	session := &Session{
		ID:        1,
		User:      User{ID: 1, Login: "testuser", Role: Role{ID: 1}},
		LoginTime: time.Now(),
		OTP:       "123456",
	}

	// Кодирование сессии в токен
	token, err := CodeSession(session, jwtKey, tokenLiveTime)
	if err != nil {
		t.Fatalf("failed to code session: %v", err)
	}

	// Декодирование токена
	decodedSession, err := GetSessionFromTkn(token, jwtKey)
	if err != nil {
		t.Fatalf("failed to get session from token: %v", err)
	}

	// Проверка корректности данных сессии
	if decodedSession.ID != session.ID {
		t.Errorf("expected session ID %d, got %d", session.ID, decodedSession.ID)
	}
	if decodedSession.User.Login != session.User.Login {
		t.Errorf("expected user login %s, got %s", session.User.Login, decodedSession.User.Login)
	}
}
