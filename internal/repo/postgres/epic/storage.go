package epic

import (
	"money/internal/repo/postgres/db"
	"money/internal/repo/postgres/stmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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

	return &Storage{stmt: stmt}, err
}
