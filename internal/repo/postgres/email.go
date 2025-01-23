package postgres

import (
	"context"
	"fmt"
	"log"

	"money/internal/core"
	"money/internal/logger"

	"go.uber.org/zap"
)

func (s *postgresStorage) AddEmail(ctx context.Context, m *core.Email) error {
	//`INSERT INTO cmn.Emails (UserID, OTP, Email) SELECT $1, $2, $3`
	err := s.preparedStatements["AddEmail"].GetContext(ctx, &m.ID, m.User.ID, m.Text, m.Email)
	if err != nil {
		return err
	}
	logger.Log.Info("Добавлено сообщение",
		zap.String(m.Email, fmt.Sprintf("С идентификатором: %d", m.ID)),
	)
	return nil
}
func (s *postgresStorage) SelectUnsendedEmails(ctx context.Context) ([]core.Email, error) {
	ms := make([]core.Email, 0)

	err := s.preparedStatements["SelectEmailForSending"].SelectContext(ctx, &ms)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *postgresStorage) UpdateEmail(ctx context.Context, c *core.Email) error {
	//UPDATE cmn.Emails SET OTP = $2, Email = $3, SendedDate = $4, Used = $5, Queued = $6 WHERE id = $1;
	_, err := s.preparedStatements["UpdateEmail"].QueryContext(ctx, c.ID, c.Text, c.Email, c.SendedDate, c.Used, c.Queued)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
