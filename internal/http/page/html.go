package page

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
)

func Execute(folder string, page string, w http.ResponseWriter, p *Page) error {
	absPath, err := filepath.Abs(filepath.Join(".", "view", folder, fmt.Sprintf("%s.html", page))) // fmt.Sprintf("./view/%s/%s.html", folder, page))
	if err != nil {
		errors.Wrap(err, "ошибка при попытке получить абсолютный путь к шаблону")
	}
	paths := getCommonPaths()
	html, err := template.ParseFiles(absPath, paths["footPath"], paths["headAuthPath"], paths["headPath"], paths["headAdminPath"], paths["headRegPath"])
	if err != nil {
		log.Println("Execute 1", err)
		return errors.Wrap(err, "ошибка при попытке разобрать шаблоны")
	}

	err = html.Execute(w, p)
	if err != nil {
		log.Println("Execute 2", err)
		return errors.Wrap(err, "ошибка при попытке запустить шаблоны")
	}
	return nil
}
