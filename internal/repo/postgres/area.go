package postgres

import (
	"context"
	"log"
	"money/internal/core"

	"github.com/pkg/errors"
)

func (s *postgresStorage) areaPrepareStmts() error {
	insertArea, err := s.DB.Preparex(`	INSERT INTO ref.Areas
										(name, levelid) 
	 									SELECT $1, $2
	 									RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertArea")
	}

	updateArea, err := s.DB.Preparex(`	UPDATE ref.Areas
										SET name = $2, levelid = $3 
	 									WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "updateArea")
	}

	deleteArea, err := s.DB.Preparex(`	DELETE FROM ref.Areas 
	 									WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "deleteArea")
	}

	selectArea, err := s.DB.Preparex(`	SELECT  
	 									a.Name, l.id AS "level.id", l.name AS "level.name" 
										FROM ref.Areas AS a
												INNER JOIN ref.Levels AS l ON l.ID = a.LevelID
	 									WHERE a.ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectArea")
	}

	selectAllAreas, err := s.DB.Preparex(`	SELECT  
	 										a.ID, a.Name, l.ID AS "level.id", l.Name AS "level.name"
											FROM ref.Areas AS a
												INNER JOIN ref.Levels AS l ON l.ID = a.LevelID;`)
	if err != nil {
		return errors.Wrap(err, "selectAllAreas")
	}
	selectAllLevelAreas, err := s.DB.Preparex(`	SELECT  
												ID, Name
										  		FROM ref.Areas
												WHERE LevelID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectAllLevelAreas")
	}

	s.preparedStatements["insertArea"] = insertArea
	s.preparedStatements["updateArea"] = updateArea
	s.preparedStatements["deleteArea"] = deleteArea
	s.preparedStatements["selectArea"] = selectArea
	s.preparedStatements["selectAllAreas"] = selectAllAreas
	s.preparedStatements["selectAllLevelAreas"] = selectAllLevelAreas
	return nil
}

func (s *postgresStorage) InsertArea(ctx context.Context, a *core.Area) error {
	log.Println("InsertArea", a.Name, a.Level.ID)
	return s.preparedStatements["insertArea"].GetContext(ctx, a, a.Name, a.Level.ID)
}
func (s *postgresStorage) UpdateArea(ctx context.Context, a *core.Area) error {
	_, err := s.preparedStatements["updateArea"].ExecContext(ctx, a.ID, a.Name, a.Level.ID)
	return err
}
func (s *postgresStorage) SelectArea(ctx context.Context, a *core.Area) error {
	return s.preparedStatements["selectArea"].GetContext(ctx, a, a.ID)
}

func (s *postgresStorage) SelectAllAreas(ctx context.Context) ([]core.Area, error) {
	a := make([]core.Area, 0)
	err := s.preparedStatements["selectAllAreas"].SelectContext(ctx, &a)
	if err != nil {

		return nil, err
	}
	return a, nil
}

func (s *postgresStorage) SelectAllLevelAreas(ctx context.Context, l *core.Level) error {
	var c []core.Area
	err := s.preparedStatements["selectLevelAreas"].SelectContext(ctx, c, l.ID)
	if err != nil {
		return err
	}
	l.PossibleAreas = c
	return nil
}

func (s *postgresStorage) DeleteArea(ctx context.Context, a *core.Area) error {
	_, err := s.preparedStatements["deleteArea"].ExecContext(ctx, a.ID)
	return err
}
