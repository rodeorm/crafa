package core

import "context"

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
