package core

type Storage struct {
	UserStorager
	MessageStorager
	SessionStorager
	DBStorager
	RoleStorager
}

type DBStorager interface {
	Ping() error
	Close() error
}
