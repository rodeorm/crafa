package sender

import (
	"context"
	"fmt"
	"money/internal/core"
	"money/internal/logger"
	"time"

	"go.uber.org/zap"
)

// Send отправляет сообщение
func (s *Sender) Send(m *core.Message) error {
	if m.Category.ID == core.MessageCategoryEmail {
		email := core.NewEmail(*m,
			core.WithHeader(s.from, m.Email),
			core.WithBody(s.domain, m.Text, m.User.ID))

		logger.Log.Info("Send",
			zap.String(fmt.Sprintf("Сендер %d получил сообщение для отправки %s", s.ID, m.Email), m.Text))

		if err := s.emailDialer.DialAndSend(email.GMS); err != nil {
			logger.Log.Error("Send",
				zap.String(fmt.Sprintf("Сендер %d не отправил сообщение", s.ID), err.Error()),
			)
			return err
		}
	}

	m.SendTime.Time = time.Now()
	m.SendTime.Valid = true
	s.msgStorager.UpdateMsg(context.TODO(), m)

	return nil
}
