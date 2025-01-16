package db

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type postgresStorage struct {
	DB *sqlx.DB
}

func InitPostgres(connectionString string) (*postgresStorage, error) {
	db := sqlx.MustConnect("pgx", connectionString)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	storage := postgresStorage{DB: db}
	return &storage, nil
}

func (s postgresStorage) CloseConnection() {
	s.DB.Close()
}
