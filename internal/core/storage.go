package core

type Storage struct {
	DBStorager
	UserStorager
	MessageStorager
	SessionStorager
	RoleStorager
	ProjectStorager
	LevelStorager
	CategoryStorager
	TeamStorager
	IterationStorager
	StatusStorager
	AreaStorager
}

type DBStorager interface {
	Ping() error
	Close() error
}
