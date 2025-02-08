package server

import (
	"context"
	"money/internal/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// gracefulShutDown реализует изящное выключение http сервера
func (s *Server) gracefulShutdown() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		// читаем из канала прерываний
		// поскольку нужно прочитать только одно прерывание,
		// можно обойтись без цикла
		<-sigint

		// создаем контекст с таймаутом
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		logger.Log.Info("Server Shutdown",
			zap.String("Изященое выклюение", "Начато"),
		)
		// получили сигнал os.Interrupt, запускаем процедуру graceful shutdown
		if err := s.srv.Shutdown(ctx); err != nil {
			// ошибки закрытия Listener
			logger.Log.Error("Server Shutdowned",
				zap.String("Изящное выключение", err.Error()),
			)
		}

		if err := s.stgs.DBStorager.Close(); err != nil {
			logger.Log.Error("Server Shutdowned",
				zap.String("Закрытие соединения с БД", err.Error()),
			)
		}
		// сообщаем основному потоку,
		// что все сетевые соединения обработаны и закрыты
		close(s.exit)
	}()
}
