package core

import (
	"context"
)

type Issue struct {
	User
	Area
	Status
	Category
	Iteration
	Epic

	Name string
	ID   int

	Comments []Comment
}

type IssueStorager interface {
	InsertIssue(context.Context, *Issue) error
	SelectIssue(context.Context, *Issue) error
	UpdateIssue(context.Context, *Issue) error
	DeleteIssue(context.Context, *Issue) error
	InsertIssueComment(context.Context, *Issue, *Comment) error
	DeleteIssueComment(context.Context, *Issue, *Comment) error
	UpdateIssueComment(context.Context, *Issue, *Comment) error
	SelectAllIssueComments(context.Context, *Issue) error
	SelectAllEpicIssues(context.Context, *Epic) ([]Issue, error)
}
