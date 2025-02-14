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
}

type DBStorager interface {
	Ping() error
	Close() error
}
