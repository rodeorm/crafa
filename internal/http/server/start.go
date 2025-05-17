package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"money/internal/cfg"
	"money/internal/http/cookie"
	"money/internal/logger"
	"money/internal/repo/cash"
	"money/internal/repo/postgres"

	_ "net/http/pprof"
)

func Start(cfg *cfg.Config, wg *sync.WaitGroup, exit chan struct{}) error {
	defer wg.Done()

	// Маршрутизаторы
	r := mux.NewRouter()
	admin := r.PathPrefix("/").Subrouter() // То, что доступно только администратору, прошедшему аутентификацию
	auth := r.PathPrefix("/").Subrouter()  // То, что доступно любому авторизованному пользователю, прошедшему аутентификацию

	// Основной сервер для обработки http-запросов
	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.RunAddress,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}
	defer srv.Close()

	ps, err := postgres.GetPostgresStorage(cfg.ConnectionString)
	if err != nil {
		return err
	}

	cm := cookie.NewCookieManager(cfg.JWTKey, cfg.TokenLiveTime)
	cash := cash.NewCashStorage()

	// Сервер с окружением
	s := &Server{srv: srv, exit: exit, cfg: cfg, ps: ps, cm: cm, cash: cash}

	configMiddlewares(r, admin, auth, s)
	configPrefixes(r)
	configPaths(r, admin, auth, s)

	logger.Log.Info("HTTP Server",
		zap.String("Порт", cfg.RunAddress),
		zap.String("БД", cfg.ConnectionString),
	)

	go http.ListenAndServe(":7070", nil) //для pprof

	s.gracefulShutdown()
	err = srv.ListenAndServe()
	if err != nil {
		logger.Log.Info("HTTP Server",
			zap.String("Порт", err.Error()),
		)
	}

	logger.Log.Info("HTTP Server",
		zap.String("Изящное выключение", "Завершено"),
	)

	return nil
}
