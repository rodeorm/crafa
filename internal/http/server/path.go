package server

import (
	"money/internal/http/middle"
	"money/internal/http/server/index"
	"money/internal/http/server/user"
	"net/http"

	"github.com/gorilla/mux"
)

func configPaths(r, admin, auth *mux.Router, s *Server) {
	r.HandleFunc("/forbidden", index.Forbidden(s))
	r.HandleFunc("/", index.Index(s))
	r.HandleFunc("/main", index.MainMenu(s))

	r.HandleFunc("/user/reg", user.RegGet).Methods(http.MethodGet)
	r.HandleFunc("/user/reg", user.RegPost(s.ps, s.cm, s.cfg.Domain)).Methods(http.MethodPost)
	r.HandleFunc("/user/wait", user.Wait(s, s.ps))
	r.HandleFunc("/user/confirm", user.ConfirmGet(s, s.ps, s.cm)).Methods(http.MethodGet)
	r.HandleFunc("/user/login", user.LoginPost(s.ps)).Methods(http.MethodPost)
	r.HandleFunc("/user/verify", user.VerifyPost(s, s.ps, s.cm, s.cfg.OTPLiveTime)).Methods(http.MethodPost)
	r.HandleFunc("/user/logout", user.LogOut)
	admin.HandleFunc("/user/list", user.ListGet(s, s.ps)).Methods(http.MethodGet)

	/*
		auth.HandleFunc("/user/update", user.UpdateGet(s, s.ps, s.ps, s.ps, s.)).Methods(http.MethodGet)
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
		admin.HandleFunc("/team/delete", s.teamDeleteGet).Methods(http.MethodGet)
		admin.HandleFunc("/team/delete", s.teamDeletePost).Methods(http.MethodPost)
		admin.HandleFunc("/team/connect", s.teamConnectPost).Methods(http.MethodPost)
		admin.HandleFunc("/team/disconnect", s.teamDisconnectGet).Methods(http.MethodGet)

		admin.HandleFunc("/category/list", s.categoryListGet).Methods(http.MethodGet)
		admin.HandleFunc("/category/create", s.categoryCreatePost).Methods(http.MethodPost)

		admin.HandleFunc("/priority/list", s.priorityListGet).Methods(http.MethodGet)
		admin.HandleFunc("/priority/create", s.priorityCreatePost).Methods(http.MethodPost)

		admin.HandleFunc("/status/list", s.statusListGet).Methods(http.MethodGet)
		admin.HandleFunc("/status/create", s.statusCreatePost).Methods(http.MethodPost)

		admin.HandleFunc("/area/list", s.areaListGet).Methods(http.MethodGet)
		admin.HandleFunc("/area/create", s.areaCreatePost).Methods(http.MethodPost)

		//admin.HandleFunc("/admin/index", s.forbidden)
	*/
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
