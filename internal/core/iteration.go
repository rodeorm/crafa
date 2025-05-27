package core

import (
	"database/sql"
)

type Iteration struct {
	Level Level
	Name  string

	ID    int
	Year  int
	Month sql.NullInt32

	Parent *Iteration
	Child  []Iteration
}
