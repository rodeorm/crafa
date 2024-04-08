package core

import (
	"time"
)

type History struct {
	Author      User
	Changer     User
	CreatedDate time.Time
	AlterDate   time.Time
}
