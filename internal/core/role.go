package core

const (
	Guest    = iota // Гость
	Admin           // Администратор
	Reg             // Зарегистрированный
	Auth            // Авторизованный
	Employee        // Сотрудник
)

type Role struct {
	ID    int
	Name  string
	Const string
}
