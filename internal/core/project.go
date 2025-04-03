package core

type Project struct {
	ID   int
	Name string

	IssueQty     int
	OpenIssueQty int

	Epics []Epic
	Users []User
}
