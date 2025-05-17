package cash

import (
	"context"
	"fmt"
	"money/internal/core"
)

type LevelCash struct {
}

func (lc *LevelCash) SelectAllLevels(ctx context.Context) ([]core.Level, error) {
	ls := make([]core.Level, 6)
	ls[0] = core.Level{ID: core.LevelProject, Name: "Проект", Const: "LevelProject"}
	ls[1] = core.Level{ID: core.LevelEpic, Name: "Эпик", Const: "LevelEpic"}
	ls[2] = core.Level{ID: core.LevelReq, Name: "Требование", Const: "LevelReq"}
	ls[3] = core.Level{ID: core.LevelIssue, Name: "Проблема", Const: "LevelIssue"}
	ls[4] = core.Level{ID: core.LevelFeature, Name: "Функциональность", Const: "LevelFeature"}
	ls[5] = core.Level{ID: core.LevelTask, Name: "Задача", Const: "LevelTask"}

	return ls, nil
}

func (lc *LevelCash) SelectLevel(ctx context.Context, l *core.Level) error {
	switch l.ID {
	case 0:
		l.Const = "LevelProject"
		l.Name = "Проект"
	case 1:
		l.Const = "LevelEpic"
		l.Name = "Эпик"
	case 2:
		l.Const = "LevelReq"
		l.Name = "Требование"
	case 3:
		l.Const = "LevelIssue"
		l.Name = "Проблема"
	case 4:
		l.Const = "LevelFeature"
		l.Name = "Функциональность"
	case 5:
		l.Const = "LevelTask"
		l.Name = "Задача"
	default:
		return fmt.Errorf("некорректный уровень")
	}

	return nil
}
