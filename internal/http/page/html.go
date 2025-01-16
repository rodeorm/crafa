package page

import (
	"fmt"
	"money/internal/core"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
)

func Execute(folder string, page string, w http.ResponseWriter, s *core.Session, p *Page) error {
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
	html, err := template.ParseFiles(absPath, footPath, headAuthPath, headPath, headAdminPath, headRegPath)
	if err != nil {
		return errors.Wrap(err, "ошибка при попытке разобрать шаблоны")
	}

	err = html.Execute(w, s)
	if err != nil {
		return errors.Wrap(err, "ошибка при попытке запустить шаблоны")
	}
	return nil
}
