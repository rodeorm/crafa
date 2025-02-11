package filler

import (
	"context"
	"log"
	"sync"
	"time"

	"money/internal/cfg"
	"money/internal/core"
	"money/internal/logger"

	"go.uber.org/zap"
)

// Filler - рабочий, заполняющий очередь сообщений
type Filler struct {
	msgStorager core.MessageStorager //16 байт. Хранилище сообщений
	queue       *core.Queue          //8 байт. Очередь сообщений
	ID          int                  //8 байт. Идентификатор воркера
	period      int                  //8 байт. Периодичность наполнения сообщений
}

func Start(config *cfg.Config, es core.MessageStorager, wg *sync.WaitGroup, exit chan struct{}) {
	// Асинхронно запускаем наполнитель очереди
	s := NewFiller(
		config.EmailQueue,
		es,
		config.QueueFillPeriod,
	)

	go s.StartFilling(exit, wg)
}

// StartFilling начинает наполнение очереди
func (f *Filler) StartFilling(exit chan struct{}, wg *sync.WaitGroup) {
	logger.Log.Info("StartFilling",
		zap.String("Филлер стартовал", "Успешно"))
	ctx := context.TODO()
	for {
		select {
		case _, ok := <-exit:
			if !ok {
				//Нет смысла ждать наполнения очереди, поэтому дефолт не жду
				logger.Log.Info("StartFilling",
					zap.String("Филлер изящно завершил дела", "Успешно"))
				wg.Done()
				return
			}
		default:
			msgs, err := f.msgStorager.SelectUnsendedMsgs(ctx)

			if err != nil {
				logger.Log.Error("StartFilling",
					zap.String("ошибка при получении сообщений к отправке", err.Error()),
				)
			}

			for _, v := range msgs {
				log.Println("Филлер пишет сообщение", v.ID)
				f.queue.Push(&v)
				log.Println("Филлер записал сообщение", v.ID)
				v.Queued = true
				f.msgStorager.UpdateMsg(ctx, &v)
				log.Println("Филлер обновил сообщение в БД", v.ID)
			}
		}
		time.Sleep(time.Duration(f.period) * time.Second)
	}
}
