package core

type Storage struct {
	UserStorager
	MessageStorager
	SessionStorager
	DBStorager
	RoleStorager
	ProjectStorager
	LevelStorager
}

type DBStorager interface {
	Ping() error
	Close() error
}
