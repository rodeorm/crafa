package core

const (
	Guest    = iota // Гость
	Admin           // Администратор
	Reg             // Зарегистрированный
	Auth            // Авторизованный
	Support         // Техническая поддержка
	Employee        // Сотрудник
)

type Role struct {
	ID   int
	Name string
}
