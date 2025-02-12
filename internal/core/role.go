package core

import (
	"context"
)

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

type RoleStorager interface {
	SelectPossibleRoles(context.Context) ([]Role, error)
	SelectRole(context.Context, *Role) error
}

type RoleCash struct {
}

func (rc *RoleCash) SelectPossibleRoles(ctx context.Context) ([]Role, error) {
	rs := make([]Role, 3)
	rs[0] = Role{ID: RoleAdmin, Name: "Администратор", Const: "RoleAdmin"}
	rs[1] = Role{ID: RoleAuth, Name: "Авторизованный пользователь", Const: "RoleAuth"}
	rs[2] = Role{ID: RoleEmployee, Name: "Сотрудник", Const: "RoleEmployee"}

	return rs, nil
}

func (rc *RoleCash) SelectRole(ctx context.Context, r *Role) error {
	switch r.ID {
	case 0:
		r.Const = "RoleGuest"
		r.Name = "Гость"
	case 1:
		r.Const = "RoleAdmin"
		r.Name = "Администратор"
	case 2:
		r.Const = "RoleReg"
		r.Name = "Зарегистрированный пользователь"
	case 3:
		r.Const = "RoleAuth"
		r.Name = "Авторизованный пользователь"
	case 4:
		r.Const = "RoleEmployee"
		r.Name = "Сотрудник"
	}

	return nil
}
