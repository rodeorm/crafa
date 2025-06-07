package user

import (
	"crafa/internal/repo/postgres/db"
	"crafa/internal/repo/postgres/stmt"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	stmt map[string]*sqlx.Stmt
	db   *sqlx.DB
}

func NewStorage(connectionString string) (*Storage, error) {
	db, err := db.GetDB(connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "New Storage. GetDB")
	}

	stmt, err := stmt.MakeStmtFromQueries(db, getQueries())
	if err != nil {
		return nil, err
	}

	return &Storage{stmt: stmt, db: db}, err
}
