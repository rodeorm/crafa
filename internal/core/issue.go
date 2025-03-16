package core

import (
	"context"
)

// Issue - это проблема или запрос на портале тех.поддержки
type Issue struct {
	Author    User
	Supporter User
	Area      Area
	Status    Status
	Category  Category
	Iteration Iteration
	Epic      Epic
	Project   Project
	Name      string
	Text      string
	ID        int
	Comments  []Comment
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
