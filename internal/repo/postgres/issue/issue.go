package issue

import (
	"context"
	"log"
	"crafa/internal/core"
	"crafa/internal/logger"

	"go.uber.org/zap"
)

func (s *Storage) InsertIssue(ctx context.Context, p *core.Issue) error {
	_, err := s.stmt["insertIssue"].ExecContext(ctx, p.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateIssue(ctx context.Context, p *core.Issue) error {
	//SET Name = $2 WHERE ID = $1
	_, err := s.stmt["updateIssue"].ExecContext(ctx, p.ID, p.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) SelectIssue(ctx context.Context, p *core.Issue) error {

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("selectIssue 1", err)
		return err
	}
	err = tx.Stmtx(s.stmt["selectIssue"]).GetContext(ctx, p, p.ID)
	if err != nil {
		log.Println("selectIssue 2", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()

}

func (s *Storage) SelectAllIssues(ctx context.Context) ([]core.Issue, error) {
	p := make([]core.Issue, 0)
	err := s.stmt["selectAllIssues"].SelectContext(ctx, &p)
	if err != nil {
		logger.Log.Error("selectAllIssues",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *Storage) SelectUserIssues(ctx context.Context, u *core.User) ([]core.Issue, error) {
	p := make([]core.Issue, 0)
	err := s.stmt["selectUserIssues"].SelectContext(ctx, &p, u.ID)
	if err != nil {
		logger.Log.Error("selectAllIssues",
			zap.Error(err))
		return nil, err

	}
	return p, nil
}

func (s *Storage) DeleteIssue(ctx context.Context, p *core.Issue) error {
	_, err := s.stmt["deleteIssue"].ExecContext(ctx, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteUserIssue(ctx context.Context, u *core.User, p *core.Issue) error {
	_, err := s.stmt["deleteUserIssue"].ExecContext(ctx, u.ID, p.ID)
	if err != nil {
		logger.Log.Error("DeleteUserIssue",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *Storage) InsertUserIssue(ctx context.Context, userID, IssueID int) error {
	_, err := s.stmt["insertUserIssue"].ExecContext(ctx, userID, IssueID)
	if err != nil {
		logger.Log.Error("InsertUserIssue",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *Storage) SelectAllIssueEpics(ctx context.Context, c *core.Issue) ([]core.Epic, error) {
	return nil, nil
}
