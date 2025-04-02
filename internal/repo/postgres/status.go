package postgres

import (
	"context"
	"money/internal/core"

	"github.com/pkg/errors"
)

func (s *PostgresStorage) statusPrepareStmts() error {
	insertStatus, err := s.DB.Preparex(`	INSERT INTO ref.Statuses
										(name, levelid) 
	 									SELECT $1, $2
	 									RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertStatus")
	}

	insertStatusHierarchy, err := s.DB.Preparex(`	INSERT INTO ref.StatusHierarchy
													(parent, child) 
	 												SELECT $1, $2;`)
	if err != nil {
		return errors.Wrap(err, "insertStatusHierarchy")
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

	selectFirstLevelStatuses, err := s.DB.Preparex(`	SELECT  
															s.ID, s.Name
										  				FROM ref.Statuses AS s
															LEFT JOIN ref.StatusHierarchy AS sh ON sh.Child = s.ID
														WHERE LevelID = $1 AND sh.Parent IS NULL;`)
	if err != nil {
		return errors.Wrap(err, "selectFirstLevelStatuses")
	}

	selectParents, err := s.DB.Preparex(`	SELECT  
														s.ID, s.Name
													  FROM ref.Statuses AS s
														INNER JOIN ref.StatusHierarchy AS sh ON sh.Child = s.ID
													WHERE LevelID = $1;`)

	if err != nil {
		return errors.Wrap(err, "selectParents")
	}
	selectPossibleParents, err := s.DB.Preparex(`	SELECT  
														s.ID, s.Name
													  FROM ref.Statuses AS s
														LEFT JOIN ref.StatusHierarchy AS sh ON sh.Child = s.ID
													WHERE LevelID = $1 AND sh.Parent IS NULL;`)

	if err != nil {
		return errors.Wrap(err, "selectPossibleParents")
	}
	selectChildren, err := s.DB.Preparex(`	SELECT  
													s.ID, s.Name
												  FROM ref.Statuses AS s
													INNER JOIN ref.StatusHierarchy AS sh ON sh.Child = s.ID
												WHERE LevelID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectChildren")
	}

	selectPossibleChildren, err := s.DB.Preparex(`	SELECT  
													s.ID, s.Name
												  FROM ref.Statuses AS s
													LEFT JOIN ref.StatusHierarchy AS sh ON sh.Child = s.ID
												WHERE LevelID = $1 AND sh.Parent IS NULL;`)
	if err != nil {
		return errors.Wrap(err, "selectPossibleChildren")
	}

	s.preparedStatements["insertStatus"] = insertStatus
	s.preparedStatements["insertStatusHierarchy"] = insertStatusHierarchy
	s.preparedStatements["updateStatus"] = updateStatus
	s.preparedStatements["deleteStatus"] = deleteStatus
	s.preparedStatements["selectStatus"] = selectStatus
	s.preparedStatements["selectAllLevelStatuses"] = selectAllLevelStatuses
	s.preparedStatements["selectAllStatuses"] = selectAllStatuses
	s.preparedStatements["selectFirstLevelStatuses"] = selectFirstLevelStatuses
	s.preparedStatements["selectStatusParents"] = selectParents
	s.preparedStatements["selectStatusPossibleParents"] = selectPossibleParents
	s.preparedStatements["selectStatusChildren"] = selectChildren
	s.preparedStatements["selectStatusPossibleChildren"] = selectPossibleChildren
	return nil
}

func (s *PostgresStorage) InsertStatus(ctx context.Context, a *core.Status) error {
	return s.preparedStatements["insertStatus"].GetContext(ctx, a, a.Name, a.Level.ID)
}

func (s *PostgresStorage) InsertStatusHierarchy(ctx context.Context, parentID, childID int) error {
	_, err := s.preparedStatements["insertStatusHierarchy"].ExecContext(ctx, parentID, childID)
	return err
}

func (s *PostgresStorage) UpdateStatus(ctx context.Context, a *core.Status) error {
	_, err := s.preparedStatements["updateStatus"].ExecContext(ctx, a.ID, a.Name, a.Level.ID)
	return err
}
func (s *PostgresStorage) SelectStatus(ctx context.Context, st *core.Status) error {
	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "SelectStatus 1")
	}

	st.Children = make([]core.Status, 0)
	st.Parents = make([]core.Status, 0)

	// Получаем основные данные статуса
	err = tx.Stmtx(s.preparedStatements["selectStatus"]).GetContext(ctx, st, st.ID)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "SelectStatus 2")
	}
	// Получаем родительские статусы
	err = tx.Stmtx(s.preparedStatements["selectParents"]).GetContext(ctx, &st.Parents, st.ID)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "SelectStatus 3")
	}
	// Получаем дочерние статусы
	err = tx.Stmtx(s.preparedStatements["selectChildren"]).GetContext(ctx, &st.Children, st.ID)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "SelectStatus 4")
	}
	return tx.Commit()
}

func (s *PostgresStorage) SelectAllStatuses(ctx context.Context) ([]core.Status, error) {
	a := make([]core.Status, 0)
	err := s.preparedStatements["selectAllStatuses"].SelectContext(ctx, &a)
	if err != nil {

		return nil, err
	}
	return a, nil
}

func (s *PostgresStorage) SelectAllLevelStatuses(ctx context.Context, l *core.Level) error {
	var c []core.Status
	err := s.preparedStatements["selectLevelStatuses"].SelectContext(ctx, c, l.ID)
	if err != nil {
		return err
	}
	l.PossibleStatuses = c
	return nil
}

func (s *PostgresStorage) SelectFirstLevelStatuses(ctx context.Context, l *core.Level) error {
	var c []core.Status
	err := s.preparedStatements["selectFirstLevelStatuses"].SelectContext(ctx, c, l.ID)
	if err != nil {
		return err
	}
	l.PossibleStatuses = c
	return nil
}

func (s *PostgresStorage) SelectPossibleParents(ctx context.Context, st *core.Status) ([]core.Status, error) {
	var parents []core.Status
	err := s.preparedStatements["selectStatusPossibleParents"].SelectContext(ctx, parents, st.ID)
	if err != nil {
		return nil, err
	}
	return parents, nil
}

func (s *PostgresStorage) SelectPossibleChildren(ctx context.Context, st *core.Status) ([]core.Status, error) {
	var parents []core.Status
	err := s.preparedStatements["selectStatusPossibleChildren"].SelectContext(ctx, parents, st.ID)
	if err != nil {
		return nil, err
	}
	return parents, nil
}

func (s *PostgresStorage) DeleteStatus(ctx context.Context, a *core.Status) error {
	_, err := s.preparedStatements["deleteStatus"].ExecContext(ctx, a.ID)
	return err
}
