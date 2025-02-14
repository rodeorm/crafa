package postgres

import (
	"money/internal/logger"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Реализация хранилища в СУБД Postgres
type postgresStorage struct {
	DB                 *sqlx.DB
	preparedStatements map[string]*sqlx.Stmt
}

// GetPostgresStorage возвращает хранилище данных в Postgres (создает, если его не было ранее)
func GetPostgresStorage(connectionString, cryptKey string) (*postgresStorage, error) {
	var (
		dbErr error
		db    *sqlx.DB
		ps    *postgresStorage
		once  sync.Once
	)
	once.Do(
		func() {
			db, dbErr = sqlx.Open("pgx", connectionString)
			db.SetMaxOpenConns(50) // Подобрать оптимальное значение
			db.SetMaxIdleConns(2)  // Подобрать оптимальное значение
			db.SetConnMaxLifetime(10 * time.Second)

			if dbErr != nil {

				return
			}
			ps = &postgresStorage{DB: db, preparedStatements: map[string]*sqlx.Stmt{}}
			dbErr = ps.prepareStmts()
		})

	if dbErr != nil {
		logger.Log.Error("GetPostgresStorage",
			zap.String("ошибка при инициализации подключения к БД", dbErr.Error()),
		)
		return nil, dbErr
	}

	return ps, nil
}

func (s postgresStorage) Close() error {
	return s.DB.Close()
}

func (s postgresStorage) Ping() error {
	return s.DB.Ping()
}
