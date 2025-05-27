package stmt

import (
	"money/internal/logger"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// MakeStmtFromQueries на основании запросов возвращает стейтменты
func MakeStmtFromQueries(db *sqlx.DB, q map[string]string) (map[string]*sqlx.Stmt, error) {
	stmts := make(map[string]*sqlx.Stmt, 0)
	for i, v := range q {
		stmt, err := db.Preparex(v)
		if err != nil {
			logger.Log.Error("MakeStmtFromQueries",
				zap.String(i, err.Error()),
			)
			return nil, errors.Wrap(err, i)
		}
		stmts[i] = stmt
	}
	return stmts, nil
}
