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
type PostgresStorage struct {
	DB                 *sqlx.DB
	preparedStatements map[string]*sqlx.Stmt
}

var (
	dbErr error
	db    *sqlx.DB
	ps    *PostgresStorage
	once  sync.Once
)

// GetPostgresStorage возвращает хранилище данных в Postgres (создает, если его не было ранее)
func GetPostgresStorage(connectionString string) (*PostgresStorage, error) {
	once.Do(
		func() {
			db, dbErr = sqlx.Open("pgx", connectionString)
			db.SetMaxOpenConns(50) // Подобрать оптимальное значение
			db.SetMaxIdleConns(2)  // Подобрать оптимальное значение
			db.SetConnMaxLifetime(10 * time.Second)

			if dbErr != nil {

				return
			}
			ps = &PostgresStorage{DB: db, preparedStatements: map[string]*sqlx.Stmt{}}
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

func (s PostgresStorage) Close() error {
	return s.DB.Close()
}

func (s PostgresStorage) Ping() error {
	return s.DB.Ping()
}
