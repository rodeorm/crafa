package crypt

import (
	"money/internal/core"
)

func GetRoleIDFromTkn(tknStr string) (int, error) {
	cl, err := GetClaims(tknStr)
	if err != nil {
		return 0, err
	}

	return cl.RoleID, nil
}

func GetSessionFromTkn(tknStr string) (*core.Session, error) {
	cl, err := GetClaims(tknStr)
	if err != nil {
		return nil, err
	}
	return &core.Session{ID: cl.SessionID, User: core.User{ID: cl.UserID, Login: cl.Login}}, nil

}
