package msg

import (
	"context"
	"log"
	"money/internal/core"
)

func (s *Storage) UpdateMsg(ctx context.Context, e *core.Message) error {
	_, err := s.stmt["updateMsg"].ExecContext(ctx, e.ID, e.Used, e.Queued, e.SendTime)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
