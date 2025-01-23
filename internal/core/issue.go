package core

type Issue struct {
	Creator   User
	Supporter User

	Name string
	ID   int
}
