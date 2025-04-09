package server

import (
	"money/internal/http/middle"
	"money/internal/http/server/area"
	"money/internal/http/server/category"
	"money/internal/http/server/index"
	"money/internal/http/server/iteration"
	"money/internal/http/server/priority"
	"money/internal/http/server/project"
	"money/internal/http/server/status"
	"money/internal/http/server/team"
	"money/internal/http/server/user"
	"net/http"

	"github.com/gorilla/mux"
)

func configPaths(r, admin, auth *mux.Router, s *Server) {
	r.HandleFunc("/forbidden", index.Forbidden(s))
	r.HandleFunc("/", index.Index(s))
	r.HandleFunc("/main", index.MainMenu(s))

	r.HandleFunc("/user/reg", user.RegGet).Methods(http.MethodGet)
	r.HandleFunc("/user/reg", user.RegPost(s.ps.User, s.cm, s.cfg.Domain)).Methods(http.MethodPost)
	r.HandleFunc("/user/wait", user.Wait(s, s.ps.User))
	r.HandleFunc("/user/confirm", user.ConfirmGet(s, s.ps.User, s.cm)).Methods(http.MethodGet)
	r.HandleFunc("/user/login", user.LoginPost(s.ps.User)).Methods(http.MethodPost)
	r.HandleFunc("/user/verify", user.VerifyPost(s, s.ps.User, s.cm, s.cfg.OTPLiveTime)).Methods(http.MethodPost)
	r.HandleFunc("/user/logout", user.LogOut)
	admin.HandleFunc("/user/list", user.ListGet(s, s.ps.User)).Methods(http.MethodGet)
	auth.HandleFunc("/user/update", user.UpdateGet(s, s.ps.User, s.ps.Project, s.ps.Team, s.cash)).Methods(http.MethodGet)
	auth.HandleFunc("/user/update", user.UpdatePost(s, s.ps.User, s.ps.Project, s.ps.Team, s.cash)).Methods(http.MethodPost)

	auth.HandleFunc("/project/list", project.ListGet(s, s.ps.Project, s.ps.Project)).Methods(http.MethodGet)
	admin.HandleFunc("/project/create", project.CreatePost(s, s.ps.Project)).Methods(http.MethodPost)
	admin.HandleFunc("/project/update", project.UpdateGet(s, s.ps.Project)).Methods(http.MethodGet)
	admin.HandleFunc("/project/update", project.UpdatePost(s, s.ps.Project)).Methods(http.MethodPost)
	admin.HandleFunc("/project/connect", project.ConnectPost(s, s.ps.Project)).Methods(http.MethodPost)
	admin.HandleFunc("/project/disconnect", project.DisconnectGet(s, s.ps.Project, s.ps.Project)).Methods(http.MethodGet)

	auth.HandleFunc("/iteration/list", iteration.ListGet(s, s.ps.Iteration, s.cash, s.ps.Team)).Methods(http.MethodGet)
	admin.HandleFunc("/iteration/create", iteration.CreatePost(s, s.ps.Iteration)).Methods(http.MethodPost)
	admin.HandleFunc("/iteration/update", iteration.UpdateGet(s, s.ps.Iteration)).Methods(http.MethodGet)
	admin.HandleFunc("/iteration/update", iteration.UpdatePost(s, s.ps.Iteration)).Methods(http.MethodPost)

	admin.HandleFunc("/category/list", category.ListGet(s, s.ps.Category, s.cash)).Methods(http.MethodGet)
	admin.HandleFunc("/category/create", category.CreatePost(s, s.ps.Category)).Methods(http.MethodPost)

	admin.HandleFunc("/team/list", team.ListGet(s, s.ps.Team)).Methods(http.MethodGet)
	admin.HandleFunc("/team/create", team.CreatePost(s, s.ps.Team)).Methods(http.MethodPost)
	admin.HandleFunc("/team/update", team.UpdateGet(s, s.ps.Team)).Methods(http.MethodGet)
	admin.HandleFunc("/team/update", team.UpdatePost(s, s.ps.Team)).Methods(http.MethodPost)
	admin.HandleFunc("/team/delete", team.DeleteGet(s, s.ps.Team)).Methods(http.MethodGet)
	admin.HandleFunc("/team/delete", team.DeletePost(s, s.ps.Team)).Methods(http.MethodPost)
	admin.HandleFunc("/team/connect", team.ConnectPost(s, s.ps.Team)).Methods(http.MethodPost)
	admin.HandleFunc("/team/disconnect", team.DisconnectGet(s, s.ps.Team)).Methods(http.MethodGet)

	admin.HandleFunc("/priority/list", priority.ListGet(s, s.cash, s.ps.Priority)).Methods(http.MethodGet)
	admin.HandleFunc("/priority/create", priority.CreatePost(s, s.ps.Priority)).Methods(http.MethodPost)

	admin.HandleFunc("/status/list", status.ListGet(s, s.ps.Status, s.cash)).Methods(http.MethodGet)
	admin.HandleFunc("/status/create", status.CreatePost(s, s.ps.Status)).Methods(http.MethodPost)

	admin.HandleFunc("/area/list", area.ListGet(s, s.ps.Area, s.cash)).Methods(http.MethodGet)
	admin.HandleFunc("/area/create", area.CreatePost(s, s.ps.Area)).Methods(http.MethodPost)

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
