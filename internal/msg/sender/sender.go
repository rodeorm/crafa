package sender

import (
	"fmt"
	"sync"
	"time"

	"money/internal/cfg"
	"money/internal/core"
	"money/internal/logger"
	"money/internal/repo/postgres"

	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

// Sender - рабочий, отправляющий сообщения
type Sender struct {
	msgStorager MessageStorager // Хранилище сообщений
	domain      string          // Домен
	from        string          // Отправитель
	fileName    string          // Имя файла вложения
	queue       QueueStorager   // Очередь сообщений
	emailDialer *gomail.Dialer  // Отправитель
	ID          int             // Идентификатор воркера
	period      int             // Периодичность отправки сообщений
}

type QueueStorager interface {
	PopWait() *core.Message
	Len() int
}

// NewSender создает новый Sender
// Каждый Sender может рассылать сообщения через свой собственный smtp сервер
func NewSender(queue QueueStorager, storage MessageStorager, id, smtpPort, prd int, smtpServer, smtpLogin, smtpPassword, from, fileName, domain string) *Sender {
	s := Sender{
		ID:          id,
		queue:       queue,
		msgStorager: storage,
		domain:      domain,
		period:      prd,
		from:        from,
		fileName:    fileName,
	}

	s.emailDialer = gomail.NewDialer(smtpServer, smtpPort, smtpLogin, smtpPassword)

	return &s
}

func Start(config *cfg.Config, wg *sync.WaitGroup, exit chan struct{}) {
	ps, _ := postgres.GetPostgresStorage(config.ConnectionString)

	for i := range config.SendWorkerCount {
		// Асинхронно запускаем email сендеры
		s := NewSender(
			config.Queue,
			ps,
			i,
			config.SMTPPort,
			config.MessageSendPeriod,
			config.SMTPServer,
			config.SMTPLogin,
			config.SMTPPass,
			config.From,
			config.File,
			config.Domain,
		)

		go s.StartSending(exit, wg)
	}
}

// StartSending начинает отправку сообщений
func (s *Sender) StartSending(exit chan struct{}, wg *sync.WaitGroup) {
	logger.Log.Info("StartSending",
		zap.String(fmt.Sprintf("Сендер %d", s.ID), "стартовал"),
	)

	var wg_w sync.WaitGroup

	for {

		select {
		case _, ok := <-exit:
			if !ok {
				wg_w.Wait()
				for s.queue.Len() != 0 { // Если в очереди еще что-то осталось, то выгребаем
					ms := s.queue.PopWait()
					err := s.Send(ms)
					if err != nil {
						logger.Log.Error("StartSending",
							zap.String(fmt.Sprintf("Сендер %d", s.ID), err.Error()),
						)
						break
					}
					logger.Log.Info("StartSending",
						zap.Int(fmt.Sprintf("Сендер %d отправил сообщение после отмены", s.ID), ms.ID),
					)
				}

				logger.Log.Info("StartSending",
					zap.String(fmt.Sprintf("Сендер %d", s.ID), "изящно завершил дела"),
				)
				wg.Done()
				return
			}
		default:

			wg_w.Add(1)

			go func() {
				ms := s.queue.PopWait()

				if ms == nil {
					wg_w.Done()
					return
				}
				err := s.Send(ms)
				if err != nil {
					logger.Log.Error("StartSending",
						zap.String(fmt.Sprintf("Сендер %d", s.ID), err.Error()),
					)
					wg_w.Done()
					return
				}
				logger.Log.Info("StartSending",
					zap.String(fmt.Sprintf("Сендер %d", s.ID), fmt.Sprintf("отправил сообщение по адресу %s", ms.Email)),
				)
				wg_w.Done()
			}()
			time.Sleep(time.Duration(s.period) * time.Second)
		}

	}
}
