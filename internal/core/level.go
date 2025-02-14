package core

import "context"

const (
	LevelProject = iota // Уровень "Проект"
	LevelEpic           // Уровень "Эпик"
	LevelReq            // Уровень "Требование"
	LevelIssue          // Уровень "Проблема"
	LevelFeature        // Уровень "Функциональность"
	LevelTask           // Уровень "Задача"
)

type Level struct {
	Name  string
	Const string
	ID    int

	PossibleCategories []Category
}

type LevelStorager interface {
	//	AddLevel(context.Context, *Level, *User) error
	//	EditLevel(context.Context, *Level, *User) error
	SelectLevel(context.Context, *Level) error
	SelectAllLevels(context.Context) ([]Level, error)
	//	DeleteLevel(context.Context, *Level, *User) error
}

type LevelCash struct {
}

func (lc *LevelCash) SelectAllLevels(ctx context.Context) ([]Level, error) {
	ls := make([]Level, 6)
	ls[0] = Level{ID: LevelProject, Name: "Проект", Const: "LevelProject"}
	ls[1] = Level{ID: LevelEpic, Name: "Эпик", Const: "LevelEpic"}
	ls[2] = Level{ID: LevelReq, Name: "Требование", Const: "LevelReq"}
	ls[3] = Level{ID: LevelIssue, Name: "Проблема", Const: "LevelIssue"}
	ls[4] = Level{ID: LevelFeature, Name: "Функциональность", Const: "LevelFeature"}
	ls[5] = Level{ID: LevelTask, Name: "Задача", Const: "LevelTask"}

	return ls, nil
}

func (lc *LevelCash) SelectLevel(ctx context.Context, l *Level) error {
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
	}

	return nil
}
