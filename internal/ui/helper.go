package ui

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/pkg/errors"

	"money/internal/cookie"
	"money/internal/core"
)

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func makeURLWithAttributes(origin string, params map[string]string) string {
	var paramPart string

	for key, value := range params {
		if value != "" {
			paramPart = paramPart + key + "=" + value + "&"
		}
	}
	return "/" + origin + "?" + paramPart
}

func (h Handler) getCurrentUserInformation(r *http.Request) (*sessionInformation, error) {
	claims, err := cookie.GetClaimsFromCookie(r)

	if err != nil {
		log.Println("getCurrentUserInformation", err)
		return &sessionInformation{}, fmt.Errorf("пользователь не найден в куках")
	}

	session, err := h.Auth.SelectActiveSession(claims.UserID, claims.SessionID)
	if err != nil {
		log.Println("getCurrentUserInformation", err)
		return &sessionInformation{}, fmt.Errorf("для пользователя нет активной сессии")
	}

	session.LastActionTime.Time = time.Now()
	session.LastActionTime.Valid = true
	h.Auth.UpdateSession(session)
	return &sessionInformation{User: core.User{ID: claims.UserID, Login: claims.Login, Role: core.Role{ID: claims.RoleID, Name: claims.RoleName}}}, nil
}

// executeHTML инкапсулирует работу с шаблонами и генерацию html
func ExecuteHTML(folder string, page string, w http.ResponseWriter, param sessionInformation) error {
	absPath, err := filepath.Abs(filepath.Join(".", "view", folder, fmt.Sprintf("%s.html", page))) // fmt.Sprintf("./view/%s/%s.html", folder, page))
	if err != nil {
		return errors.Wrap(err, "ошибка при получении абсолютного пути для шаблонов")
	}
	footPath, err := filepath.Abs(filepath.Join(".", "view", "common", "footer.html"))
	if err != nil {
		return errors.Wrap(err, "ошибка при получении абсолютного пути для шаблонов")
	}
	headPath, err := filepath.Abs(filepath.Join(".", "view", "common", "header.html"))
	if err != nil {
		return errors.Wrap(err, "ошибка при получении абсолютного пути для шаблонов")
	}
	headAuthPath, err := filepath.Abs(filepath.Join(".", "view", "common", "headerAuth.html"))
	if err != nil {
		return errors.Wrap(err, "ошибка при получении абсолютного пути для шаблонов")
	}
	headAdminPath, err := filepath.Abs(filepath.Join(".", "view", "common", "headerAdmin.html"))
	if err != nil {
		return errors.Wrap(err, "ошибка при получении абсолютного пути для шаблонов")
	}
	headRegPath, err := filepath.Abs(filepath.Join(".", "view", "common", "headerReg.html"))
	if err != nil {
		return errors.Wrap(err, "ошибка при получении абсолютного пути для шаблонов")
	}
	var html *template.Template
	html, _ = template.ParseFiles(absPath, footPath, headAuthPath, headPath, headAdminPath, headRegPath)
	err = html.Execute(w, param)
	if err != nil {
		return errors.Wrap(err, "ошибка при попытке разобрать шаблоны")
	}
	return nil
}
