package ui

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"money/internal/core"
)

func StartUI(Auth AuthRepo, wp WorkPlaceRepo, runAddress string) error {

	main := mux.NewRouter()
	srv := &http.Server{
		Handler:      main,
		Addr:         runAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	h := Handler{Auth: Auth, Workplace: wp}

	//	То, что доступно только администратору, прошедшему аутентификацию
	admin := main.PathPrefix("/").Subrouter()
	admin.Use(h.authMiddleware)
	admin.Use(h.adminMiddleware)

	//	То, что доступно любому авторизованному пользователю, прошедшему аутентификацию
	auth := main.PathPrefix("/").Subrouter()
	auth.Use(h.authMiddleware)
	auth.Use(h.logMiddleware)

	//	Обработка статичных файлов
	staticDir := "/static/"
	staticAdminDir := "/admin/static/"

	main.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./"+staticDir))))
	admin.PathPrefix(staticAdminDir).Handler(http.StripPrefix(staticAdminDir, http.FileServer(http.Dir("./"+staticDir))))

	//	Запрет на доступ
	main.HandleFunc("/forbidden", h.forbidden)
	//	Стартовая страница
	main.HandleFunc("/", h.index)
	auth.HandleFunc("/user/logout", h.logOutPost)
	main.HandleFunc("/user/login", h.loginPost)
	main.HandleFunc("/user/reg", h.regGet).Methods(http.MethodGet)
	main.HandleFunc("/user/reg", h.regPost).Methods(http.MethodPost)

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

type Handler struct {
	Workplace WorkPlaceRepo // Реализация источника данных Omni
	Auth      AuthRepo      // Реализация источника данных подсистемы авторизации и аутентификации
}

// sessionInformation реализует передачу информации о текущей сессии, а также выбранном лицевом счете в шаблоны
type sessionInformation struct {
	User core.User

	Attribute    interface{}
	AttributeMap map[string]interface{}

	AccountID int
	Signal    string
}
