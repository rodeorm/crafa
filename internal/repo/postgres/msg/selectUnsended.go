package msg

import (
	"context"
	"log"
	"money/internal/core"
)

func (s *Storage) SelectUnsendedMsgs(ctx context.Context) ([]core.Message, error) {
	ms := make([]core.Message, 0)

	err := s.stmt["selectUnsendedMsgs"].SelectContext(ctx, &ms)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return ms, nil
}
