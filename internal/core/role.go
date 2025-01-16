package core

const (
	Guest = iota // Гость
	Admin        // Администратор
	Reg          // Зарегистрированный
	Auth         // Авторизованный
)

type Role struct {
	ID   int
	Name string
}
