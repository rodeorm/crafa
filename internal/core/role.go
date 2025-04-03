package core

const (
	RoleGuest    = iota // Гость
	RoleAdmin           // Администратор
	RoleReg             // Зарегистрированный
	RoleAuth            // Авторизованный
	RoleEmployee        // Сотрудник
)

type Role struct {
	ID    int
	Name  string
	Const string
}
