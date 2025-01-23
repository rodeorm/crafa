package core

type Project struct {
	ID   int
	Name string

	Manager    User
	Supporter  User
	Maintainer User
}

type ProjectStorager interface {
}
