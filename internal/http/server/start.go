package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"money/internal/cfg"
	"money/internal/logger"
)

func Start(a cfg.ServerConfig, wg *sync.WaitGroup, exit chan struct{}) error {
	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         a.RunAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	s := Server{srv: srv, exit: exit, cfg: &a}

	s.gracefulShutDown()
	srv.ListenAndServe()
	// ждём завершения процедуры graceful shutdown
	<-s.exit
	// получили оповещение о завершении
	// например закрыть соединение с базой данных,
	// закрыть открытые файлы
	logger.Log.Info("http server shutdowned",
		zap.String("Завершили изящное выключение", s.cfg.RunAddress),
	)
	return nil
}

func (s *Server) gracefulShutDown() {
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

		// получили сигнал os.Interrupt, запускаем процедуру graceful shutdown
		if err := s.srv.Shutdown(ctx); err != nil {
			// ошибки закрытия Listener
			logger.Log.Error("Server Shutdowned",
				zap.String("Ошибка при изящном выключении", "Сервер без https"),
			)
		}
		// сообщаем основному потоку,
		// что все сетевые соединения обработаны и закрыты
		logger.Log.Info("Server Shutdown",
			zap.String("Начали изящное выключение", ""),
		)
		close(s.exit)
	}()
}
