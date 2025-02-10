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
}

type LevelStorager interface {
	//	AddLevel(context.Context, *Level, *User) error
	//	EditLevel(context.Context, *Level, *User) error
	SelectLevel(context.Context, *Level, *User) error
	SelectAllLevels(context.Context, *User) ([]Level, error)
	//	DeleteLevel(context.Context, *Level, *User) error
}

type Catergory struct {
	Level
	Name string
	ID   int
}

type CategoryStorager interface {
	AddCategory(context.Context, *Category, *User) error
	EditCategory(context.Context, *Category, *User) error
	SelectCategory(context.Context, *Category, *User) error
	SelectAllCategories(context.Context, *User) ([]Category, error)
	DeleteCategory(context.Context, *Category, *User) error
}

type Area struct {
	Level
	Team
	ID int
}

type AreaStorager interface {
	AddArea(context.Context, *Area, *User) error
	EditArea(context.Context, *Area, *User) error
	SelectArea(context.Context, *Area, *User) error
	SelectAllAreas(context.Context, *User) ([]Area, error)
	DeleteArea(context.Context, *Area, *User) error
}

type Status struct {
	Level
	ID int
}

type StatusStorager interface {
	AddStatus(context.Context, *Status, *User) error
	EditStatus(context.Context, *Status, *User) error
	SelectStatus(context.Context, *Status, *User) error
	SelectAllStatuses(context.Context, *User) ([]Status, error)
	DeleteStatus(context.Context, *Status, *User) error
}

type Iteration struct {
	Level
	Name   string
	ID     int
	Year   int
	Month  int
	Parent *Iteration
	Child  *Iteration
}

type IterationStorager interface {
	AddIteration(context.Context, *Iteration, *User) error
	EditIteration(context.Context, *Iteration, *User) error
	SelectIteration(context.Context, *Iteration, *User) error
	SelectAllIterations(context.Context, *User) ([]Iteration, error)
	DeleteIteration(context.Context, *Iteration, *User) error
}
