package core

import (
	"database/sql"
)

type Session struct {
	ID   int
	User *User

	LogInTime      sql.NullTime
	LogOutTime     sql.NullTime
	LastActionTime sql.NullTime

	Device string
	IP     string
}
