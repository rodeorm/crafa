package core

// User - это сесияя для пользователя
type User struct {
	ID       int    `db:"user.id"`
	Login    string // Логин
	Password string // Пароль

	FamilyName string // Фамилия
	Name       string // Имя
	PatronName string // Отчество

	Email string // Адрес электронной почты
	Phone string // Телефон

	Role Role
}
