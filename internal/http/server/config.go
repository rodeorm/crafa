package server

import (
	"money/internal/http/middle"
	"net/http"

	"github.com/gorilla/mux"
)

func configPaths(r, admin, auth *mux.Router, s *Server) {
	r.HandleFunc("/forbidden", s.forbidden)
	r.HandleFunc("/", s.index)
	r.HandleFunc("/user/reg", s.regGet).Methods(http.MethodGet)
	r.HandleFunc("/user/reg", s.regPost).Methods(http.MethodPost)
	r.HandleFunc("/user/send", s.send)
	r.HandleFunc("/user/confirm", s.confirmGet).Methods(http.MethodGet)
	r.HandleFunc("/user/login", s.loginPost).Methods(http.MethodPost)
	r.HandleFunc("/user/verify", s.verifyPost).Methods(http.MethodPost)
	r.HandleFunc("/user/logout", s.logOut)

	admin.HandleFunc("/user/list", s.userListGet).Methods(http.MethodGet)
	auth.HandleFunc("/user/update", s.userUpdateGet).Methods(http.MethodGet)
	auth.HandleFunc("/user/update", s.userUpdatePost).Methods(http.MethodPost)

	auth.HandleFunc("/project/list", s.projectListGet).Methods(http.MethodGet)
	auth.HandleFunc("/project/create", s.projectCreatePost).Methods(http.MethodPost)
	auth.HandleFunc("/project/update", s.projectUpdateGet).Methods(http.MethodGet)
	auth.HandleFunc("/project/update", s.projectUpdatePost).Methods(http.MethodPost)
	admin.HandleFunc("/project/connect", s.projectConnectPost).Methods(http.MethodPost)

	//admin.HandleFunc("/admin/index", s.forbidden)
	r.HandleFunc("/main", s.main)
}

func configMiddlewares(r, admin, auth *mux.Router, s *Server) {
	r.Use(middle.WithLog)
	admin.Use(middle.WithAdmin(s.cfg.JWTKey), middle.WithLog)
	auth.Use(middle.WithAuth(s.cfg.JWTKey), middle.WithLog)
}

func configPrefixes(r *mux.Router) {
	// Обработка статичных файлов
	staticDir := "/static/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./"+staticDir))))
}
