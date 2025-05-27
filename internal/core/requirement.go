package core

type Requirement struct {
	ID   int
	Name string
	Text string

	Author      User
	Executioner User

	Features []Feature
}
