package core

// Область
type Area struct {
	Level  Level
	ID     int
	Name   string
	Parent *Area
	Child  []Area
}
