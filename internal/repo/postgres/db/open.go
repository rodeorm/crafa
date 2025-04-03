package db

import (
	"money/internal/logger"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	dbErr error
	db    *sqlx.DB
	once  sync.Once
)

// GetDB возвращает подключение к Postgres (создает, если его не было ранее)
func GetDB(connectionString string) (*sqlx.DB, error) {
	once.Do(
		func() {
			db, dbErr = sqlx.Open("pgx", connectionString)
			db.SetMaxOpenConns(50) // Подобрать оптимальное значение
			db.SetMaxIdleConns(2)  // Подобрать оптимальное значение
			db.SetConnMaxLifetime(10 * time.Second)

			if dbErr != nil {
				return
			}
		})

	if dbErr != nil {
		logger.Log.Error("GetDB",
			zap.String("ошибка при инициализации подключения к БД", dbErr.Error()),
		)
		return nil, dbErr
	}

	return db, nil
}
