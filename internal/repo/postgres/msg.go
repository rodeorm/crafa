package postgres

import (
	"context"
	"log"

	"money/internal/core"
)

func (s *postgresStorage) SelectUnsendedMsgs(ctx context.Context) ([]core.Message, error) {
	ms := make([]core.Message, 0)

	err := s.preparedStatements["selectUnsendedMsgs"].SelectContext(ctx, &ms)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return ms, nil
}

func (s *postgresStorage) UpdateMsg(ctx context.Context, e *core.Message) error {
	//Used = $2, Queued = $3, SendTime = $4 WHERE id = $1;
	_, err := s.preparedStatements["updateMsg"].QueryContext(ctx, e.ID, e.Used, e.Queued, e.SendedDate)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
