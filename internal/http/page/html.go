package page

import (
	"fmt"
	"log"
	"money/internal/logger"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Execute(folder string, page string, w http.ResponseWriter, p *Page) error {
	absPath, err := filepath.Abs(filepath.Join(".", "view", folder, fmt.Sprintf("%s.html", page))) // fmt.Sprintf("./view/%s/%s.html", folder, page))
	if err != nil {
		logger.Log.Error("Execute",
			zap.String("absPath", err.Error()),
		)
		return err
	}
	footPath, err := filepath.Abs(filepath.Join(".", "view", "common", "footer.html"))
	if err != nil {
		logger.Log.Error("Execute",
			zap.String("footPath", err.Error()),
		)
		return err
	}
	headPath, err := filepath.Abs(filepath.Join(".", "view", "common", "header.html"))
	if err != nil {
		logger.Log.Error("Execute",
			zap.String("headPath", err.Error()),
		)
		return err
	}
	headAuthPath, err := filepath.Abs(filepath.Join(".", "view", "common", "headerAuth.html"))
	if err != nil {
		logger.Log.Error("Execute",
			zap.String("headAuthPath", err.Error()),
		)
		return err
	}
	headAdminPath, err := filepath.Abs(filepath.Join(".", "view", "common", "headerAdmin.html"))
	if err != nil {
		logger.Log.Error("Execute",
			zap.String("headAdminPath", err.Error()),
		)
		return err
	}
	headRegPath, err := filepath.Abs(filepath.Join(".", "view", "common", "headerReg.html"))
	if err != nil {
		logger.Log.Error("Execute",
			zap.String("headRegPath", err.Error()),
		)
		return err
	}

	html, err := template.ParseFiles(absPath, footPath, headAuthPath, headPath, headAdminPath, headRegPath)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "ошибка при попытке разобрать шаблоны")
	}

	err = html.Execute(w, p)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "ошибка при попытке запустить шаблоны")
	}
	return nil
}
