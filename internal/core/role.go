package core

type Role struct {
	ID     int
	Name   string
	Access []Section
	Denied []Section
}

type Section struct {
	ID      int
	Name    string
	BaseUrl string
}
