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
	admin.HandleFunc("/project/create", s.projectCreatePost).Methods(http.MethodPost)
	admin.HandleFunc("/project/update", s.projectUpdateGet).Methods(http.MethodGet)
	admin.HandleFunc("/project/update", s.projectUpdatePost).Methods(http.MethodPost)
	admin.HandleFunc("/project/connect", s.projectConnectPost).Methods(http.MethodPost)
	admin.HandleFunc("/project/disconnect", s.projectDisconnectGet).Methods(http.MethodGet)

	auth.HandleFunc("/iteration/list", s.iterationListGet).Methods(http.MethodGet)
	admin.HandleFunc("/iteration/create", s.iterationCreatePost).Methods(http.MethodPost)
	admin.HandleFunc("/iteration/update", s.iterationUpdateGet).Methods(http.MethodGet)
	admin.HandleFunc("/iteration/update", s.iterationUpdatePost).Methods(http.MethodPost)

	admin.HandleFunc("/category/list", s.categoryListGet).Methods(http.MethodGet)
	admin.HandleFunc("/category/create", s.categoryCreatePost).Methods(http.MethodPost)

	admin.HandleFunc("/team/list", s.teamListGet).Methods(http.MethodGet)
	admin.HandleFunc("/team/create", s.teamCreatePost).Methods(http.MethodPost)
	admin.HandleFunc("/team/update", s.teamUpdateGet).Methods(http.MethodGet)
	admin.HandleFunc("/team/update", s.teamUpdatePost).Methods(http.MethodPost)
	admin.HandleFunc("/team/connect", s.teamConnectPost).Methods(http.MethodPost)
	admin.HandleFunc("/team/disconnect", s.teamDisconnectGet).Methods(http.MethodGet)

	admin.HandleFunc("/category/list", s.categoryListGet).Methods(http.MethodGet)
	admin.HandleFunc("/category/add", s.categoryCreatePost).Methods(http.MethodPost)

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
