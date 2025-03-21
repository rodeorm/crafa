package postgres

import (
	"context"
	"log"
	"money/internal/core"

	"github.com/pkg/errors"
)

func (s *postgresStorage) statusPrepareStmts() error {
	insertStatus, err := s.DB.Preparex(`	INSERT INTO ref.Statuses
										(name, levelid) 
	 									SELECT $1, $2
	 									RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertStatus")
	}

	updateStatus, err := s.DB.Preparex(`	UPDATE ref.Statuses
										SET name = $2, levelid = $3 
	 									WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "updateStatus")
	}

	deleteStatus, err := s.DB.Preparex(`	DELETE FROM ref.Statuses 
	 									WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "deleteStatus")
	}

	selectStatus, err := s.DB.Preparex(`	SELECT  
	 									a.Name, l.id AS "level.id", l.name AS "level.name" 
										FROM ref.Statuses AS a
												INNER JOIN ref.Levels AS l ON l.ID = a.LevelID
	 									WHERE a.ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectStatus")
	}

	selectAllStatuses, err := s.DB.Preparex(`	SELECT  
	 										a.ID, a.Name, l.ID AS "level.id", l.Name AS "level.name"
											FROM ref.Statuses AS a
												INNER JOIN ref.Levels AS l ON l.ID = a.LevelID;`)
	if err != nil {
		return errors.Wrap(err, "selectAllStatuses")
	}
	selectAllLevelStatuses, err := s.DB.Preparex(`	SELECT  
												ID, Name
										  		FROM ref.Statuses
												WHERE LevelID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectAllLevelStatuses")
	}

	s.preparedStatements["insertStatus"] = insertStatus
	s.preparedStatements["updateStatus"] = updateStatus
	s.preparedStatements["deleteStatus"] = deleteStatus
	s.preparedStatements["selectStatus"] = selectStatus
	s.preparedStatements["selectAllStatuses"] = selectAllStatuses
	s.preparedStatements["selectAllLevelStatuses"] = selectAllLevelStatuses
	return nil
}

func (s *postgresStorage) InsertStatus(ctx context.Context, a *core.Status) error {
	log.Println("InsertStatus", a.Name, a.Level.ID)
	return s.preparedStatements["insertStatus"].GetContext(ctx, a, a.Name, a.Level.ID)
}
func (s *postgresStorage) UpdateStatus(ctx context.Context, a *core.Status) error {
	_, err := s.preparedStatements["updateStatus"].ExecContext(ctx, a.ID, a.Name, a.Level.ID)
	return err
}
func (s *postgresStorage) SelectStatus(ctx context.Context, a *core.Status) error {
	return s.preparedStatements["selectStatus"].GetContext(ctx, a, a.ID)
}

func (s *postgresStorage) SelectAllStatuses(ctx context.Context) ([]core.Status, error) {
	a := make([]core.Status, 0)
	err := s.preparedStatements["selectAllStatuses"].SelectContext(ctx, &a)
	if err != nil {

		return nil, err
	}
	return a, nil
}

func (s *postgresStorage) SelectAllLevelStatuses(ctx context.Context, l *core.Level) error {
	var c []core.Status
	err := s.preparedStatements["selectLevelStatuses"].SelectContext(ctx, c, l.ID)
	if err != nil {
		return err
	}
	l.PossibleStatuses = c
	return nil
}

func (s *postgresStorage) DeleteStatus(ctx context.Context, a *core.Status) error {
	_, err := s.preparedStatements["deleteStatus"].ExecContext(ctx, a.ID)
	return err
}
