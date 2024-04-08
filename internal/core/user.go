package core

import (
	"database/sql"
	"time"
)

/*
Учетная запись	Хранимая в компьютерной системе совокупность данных о пользователе,
необходимая для его опознавания (аутентификации) и предоставления доступа к его личным данным и настройкам.
Формируется в результате регистрации пользователя.
*/
type User struct {
	ID         int
	Login      string // Логин
	Surname    string // Фамилия
	Name       string // Имя
	Patronymic string // Отчество
	Email      string // Адрес электронной почты
	Phone      string // Телефон
	Password   string // Пароль
	Role       Role
}

type Notification struct {
	UserID      int
	Header      string
	Message     string
	CreatedDate time.Time
	ReadDite    sql.NullTime
}
