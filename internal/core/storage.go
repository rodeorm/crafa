package core

type Storage struct {
	UserStorager
	MessageStorager
	SessionStorager
	DBStorager
	RoleStorager
	ProjectStorager
	LevelStorager
	CategoryStorager
	TeamStorager
	IterationStorager
	StatusStorager
}

type DBStorager interface {
	Ping() error
	Close() error
}
