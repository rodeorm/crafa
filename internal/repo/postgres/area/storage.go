package area

import (
	"crafa/internal/logger"
	"crafa/internal/repo/postgres/db"
	"crafa/internal/repo/postgres/stmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	stmt map[string]*sqlx.Stmt
	db   *sqlx.DB
}

func NewStorage(connectionString string) (*Storage, error) {

	db, err := db.GetDB(connectionString)

	if err != nil {
		logger.Log.Error("NewStorage",
			zap.Error(err),
		)
		return nil, errors.Wrap(err, "New Storage. GetDB")
	}

	stmt, err := stmt.MakeStmtFromQueries(db, getQueries())
	if err != nil {
		logger.Log.Error("NewStorage",
			zap.Error(err),
		)
		return nil, err
	}
	return &Storage{stmt: stmt, db: db}, err
}
