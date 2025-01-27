package core

type Storage struct {
	UserStorager
	MessageStorager
	SessionStorager
	DBStorager
}

type DBStorager interface {
	Ping() error
	Close() error
}
