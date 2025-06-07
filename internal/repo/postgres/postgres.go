package postgres

import (
	"crafa/internal/logger"
	"crafa/internal/repo/postgres/area"
	"crafa/internal/repo/postgres/category"
	"crafa/internal/repo/postgres/iteration"
	"crafa/internal/repo/postgres/msg"
	"crafa/internal/repo/postgres/priority"
	"crafa/internal/repo/postgres/project"
	"crafa/internal/repo/postgres/status"
	"crafa/internal/repo/postgres/team"
	"crafa/internal/repo/postgres/user"
	"sync"

	"go.uber.org/zap"
)

// Реализация хранилища в СУБД Postgres
type PostgresStorage struct {
	Area      *area.Storage
	Category  *category.Storage
	Iteration *iteration.Storage
	Msg       *msg.Storage
	Priority  *priority.Storage
	User      *user.Storage
	Team      *team.Storage
	Status    *status.Storage
	Project   *project.Storage
}

var (
	dbErr error
	ps    *PostgresStorage
	once  sync.Once
)

// GetPostgresStorage возвращает хранилище данных в Postgres (создает, если его не было ранее)
func GetPostgresStorage(cs string) (*PostgresStorage, error) {

	once.Do(func() {
		ps = &PostgresStorage{}
		ps.Area, dbErr = area.NewStorage(cs)
		if dbErr != nil {
			return
		}
		ps.Category, dbErr = category.NewStorage(cs)
		if dbErr != nil {
			return
		}
		ps.Iteration, dbErr = iteration.NewStorage(cs)
		if dbErr != nil {
			return
		}
		ps.Msg, dbErr = msg.NewStorage(cs)
		if dbErr != nil {
			return
		}
		ps.Priority, dbErr = priority.NewStorage(cs)
		if dbErr != nil {
			return
		}
		ps.User, dbErr = user.NewStorage(cs)
		if dbErr != nil {
			return
		}
		ps.Team, dbErr = team.NewStorage(cs)
		if dbErr != nil {
			return
		}
		ps.Status, dbErr = status.NewStorage(cs)
		if dbErr != nil {
			return
		}
		ps.Project, dbErr = project.NewStorage(cs)
		if dbErr != nil {
			return
		}

	})

	if dbErr != nil {
		logger.Log.Error("GetPostgresStorage",
			zap.String("ошибка при инициализации подключения к БД", dbErr.Error()),
		)
		return nil, dbErr
	}
	return ps, nil
}
