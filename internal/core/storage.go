package core

type Storage struct {
	UserStorager
	EmailStorager
	SessionStorager
	DBStorager
}

type DBStorager interface {
	Ping() error
	Close() error
}
