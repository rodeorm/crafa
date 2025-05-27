package core

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
	PossibleAreas      []Area
	PossibleStatuses   []Status
	PossiblePriorities []Priority
}
