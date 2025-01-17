package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"money/internal/cfg"
	"money/internal/http/middle"
	"money/internal/logger"
)

func Start(cfg *cfg.ServerConfig, wg *sync.WaitGroup, exit chan struct{}) error {
	defer wg.Done()
	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.RunAddress,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}
	s := Server{srv: srv, exit: exit, cfg: cfg}
	r.Use(middle.WithLog)

	// То, что доступно только администратору, прошедшему аутентификацию
	admin := r.PathPrefix("/").Subrouter()
	admin.Use(middle.WithAdmin, middle.WithAuth, middle.WithLog)

	// То, что доступно любому авторизованному пользователю, прошедшему аутентификацию
	auth := r.PathPrefix("/").Subrouter()
	auth.Use(middle.WithAuth, middle.WithLog)

	// Обработка статичных файлов
	staticDir := "/static/"
	staticUserDir := "/user/static/"
	staticAdminDir := "/admin/static/"

	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./"+staticDir))))
	r.PathPrefix(staticUserDir).Handler(http.StripPrefix(staticUserDir, http.FileServer(http.Dir("./"+staticDir))))
	admin.PathPrefix(staticAdminDir).Handler(http.StripPrefix(staticAdminDir, http.FileServer(http.Dir("./"+staticDir))))

	//	Запрет на доступ
	r.HandleFunc("/forbidden", s.forbidden)
	//	Стартовая страница
	r.HandleFunc("/", s.index)
	s.gracefulShutDown()
	err := srv.ListenAndServe()
	if err != nil {
		logger.Log.Info("ListenAndServe",
			zap.String("Потенциальная ошибка", err.Error()),
		)
	}

	<-s.exit // получили оповещение о необходимости завершить работу
	// TODO: закрыть соединение с базой данных,
	// TODO: закрыть открытые файлы
	logger.Log.Info("http server shutdowned",
		zap.String("Изящное выключение", "Завершено"),
	)

	return nil
}
