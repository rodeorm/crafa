package postgres

import (
	"money/internal/logger"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Реализация хранилища в СУБД Postgres
type postgresStorage struct {
	DB                 *sqlx.DB              // 8 байт (только указатель). Драйвер подключения к СУБД
	preparedStatements map[string]*sqlx.Stmt //8 байт (только указатель)
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
			if dbErr != nil {

				return
			}
			ps = &postgresStorage{DB: db, preparedStatements: map[string]*sqlx.Stmt{}}
			dbErr = ps.prepareStatements()
		})

	if dbErr != nil {
		logger.Log.Error("GetPostgresStorage",
			zap.String("ошибка при инициализации подключения к БД", dbErr.Error()),
		)
		return nil, dbErr
	}

	return ps, nil
}

func (s postgresStorage) CloseConnection() {
	s.DB.Close()
}
