package core

import "context"

//Category
type Category struct {
	Level Level
	Name  string
	ID    int
}

type CategoryStorager interface {
	InsertCategory(ctx context.Context, c *Category) error
	UpdateCategory(ctx context.Context, c *Category) error
	SelectCategory(ctx context.Context, c *Category) error
	SelectAllCategories(ctx context.Context) ([]Category, error)
	SelectAllLevelCategories(ctx context.Context, l *Level) error
	DeleteCategory(ctx context.Context, c *Category) error
}
