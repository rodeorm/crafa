package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"money/internal/cfg"
	"money/internal/core"
	"money/internal/http/middle"
	"money/internal/logger"
)

func Start(cfg *cfg.Config, stgs *core.Storage, wg *sync.WaitGroup, exit chan struct{}) error {
	defer wg.Done()

	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.RunAddress,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}
	defer srv.Close()
	s := Server{srv: srv, exit: exit, cfg: cfg, stgs: stgs}
	r.Use(middle.WithLog)

	// То, что доступно только администратору, прошедшему аутентификацию
	admin := r.PathPrefix("/").Subrouter()
	admin.Use(middle.WithAdmin(s.cfg.JWTKey), middle.WithLog)

	// То, что доступно любому авторизованному пользователю, прошедшему аутентификацию
	auth := r.PathPrefix("/").Subrouter()
	auth.Use(middle.WithAuth(s.cfg.JWTKey), middle.WithLog)

	// Обработка статичных файлов
	staticDir := "/static/"
	staticUserDir := "/user/static/"
	staticAdminDir := "/admin/static/"

	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./"+staticDir))))
	r.PathPrefix(staticUserDir).Handler(http.StripPrefix(staticUserDir, http.FileServer(http.Dir("./"+staticDir))))
	admin.PathPrefix(staticAdminDir).Handler(http.StripPrefix(staticAdminDir, http.FileServer(http.Dir("./"+staticDir))))

	r.HandleFunc("/forbidden", s.forbidden)
	r.HandleFunc("/", s.index).Methods(http.MethodGet)

	r.HandleFunc("/user/reg", s.regGet).Methods(http.MethodGet)
	r.HandleFunc("/user/reg", s.regPost).Methods(http.MethodPost)
	r.HandleFunc("/user/send", s.send)
	r.HandleFunc("/user/confirm", s.confirmGet).Methods(http.MethodGet)
	r.HandleFunc("/user/login", s.loginPost).Methods(http.MethodPost)
	r.HandleFunc("/user/verify", s.verifyPost).Methods(http.MethodPost)
	r.HandleFunc("/user/logout", s.logOut)
	logger.Log.Info("HTTP Server",
		zap.String("Порт", cfg.RunAddress),
		zap.String("БД", cfg.ConnectionString),
	)
	s.gracefulShutdown()
	err := srv.ListenAndServe()
	if err != nil {
		logger.Log.Info("HTTP Server",
			zap.String("Порт", err.Error()),
		)
	}

	<-s.exit // получили оповещение о необходимости завершить работу

	// TODO: закрыть соединение с базой данных
	// s.stgs.DBStorager.Close()
	// TODO: закрыть открытые файлы

	logger.Log.Info("HTTP Server",
		zap.String("Изящное выключение", "Завершено"),
	)

	return nil
}
