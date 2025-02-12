package core

type Storage struct {
	UserStorager
	MessageStorager
	SessionStorager
	DBStorager
	RoleStorager
	ProjectStorager
}

type DBStorager interface {
	Ping() error
	Close() error
}
