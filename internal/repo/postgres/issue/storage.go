package issue

import (
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	stmt map[string]*sqlx.Stmt
}

/*
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
*/
