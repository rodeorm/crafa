package core

import (
	"database/sql"
	"time"
)

type Comment struct {
	User

	CreateTime time.Time
	UpdateTime sql.NullTime

	Text  string
	Files []File
}
