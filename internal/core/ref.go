package core

type Level struct {
	Name  string
	Const string
	ID    int
}

type Catergory struct {
	Level
	Name string
	ID   int
}

type Area struct {
	Level
	Team
	ID int
}

type Status struct {
	Level
	ID int
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
