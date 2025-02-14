package core

import "context"

type Epic struct {
	ID   int
	Name string
}

type EpicStorager interface {
	InsertEpic(context.Context, *Epic) error
	SelectEpic(context.Context, *Epic) error
	UpdateEpic(context.Context, *Epic) error
	DeleteEpic(context.Context, *Epic) error
	InsertEpicComment(context.Context, *Epic, *Comment) error
	DeleteEpicComment(context.Context, *Epic, *Comment) error
	UpdateEpicComment(context.Context, *Epic, *Comment) error
	SelectAllEpicIssues(context.Context, *Epic) ([]Issue, error)
	SelectAllEpicFeautures(context.Context, *Epic) ([]Feature, error)
}
