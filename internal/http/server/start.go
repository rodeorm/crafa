package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"money/internal/http/middle"
)

func Start(runAddress string) error {

	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         runAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	s := Server{srv: srv}

	//	То, что доступно только администратору, прошедшему аутентификацию
	admin := r.PathPrefix("/").Subrouter()
	admin.Use(middle.WithAuth, middle.WithAdmin)

	//	То, что доступно любому авторизованному пользователю, прошедшему аутентификацию
	auth := r.PathPrefix("/").Subrouter()
	auth.Use(middle.WithAuth, middle.WithLog)

	//	Обработка статичных файлов
	staticDir := "/static/"
	staticAdminDir := "/admin/static/"

	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./"+staticDir))))
	admin.PathPrefix(staticAdminDir).Handler(http.StripPrefix(staticAdminDir, http.FileServer(http.Dir("./"+staticDir))))

	//	Запрет на доступ
	r.HandleFunc("/forbidden", s.forbidden)
	//	Стартовая страница
	r.HandleFunc("/", s.index)
	auth.HandleFunc("/user/logout", s.logOutPost)
	r.HandleFunc("/user/login", s.loginPost)
	r.HandleFunc("/user/reg", s.regGet).Methods(http.MethodGet)
	r.HandleFunc("/user/reg", s.regPost).Methods(http.MethodPost)

	log.Fatal(srv.ListenAndServe())
	/*	Если появится https и SSL сертификат
		var (
			certFile string - относительный путь к сертификату
			keyFile  string - относительный путь к ключу
		)
		log.Fatal(srv.ListenAndServeTLS(certFile, keyFile))
	*/
	return nil
}
