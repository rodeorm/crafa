package cash

import (
	"context"
	"crafa/internal/core"
)

type RoleCash struct {
}

func (rc *RoleCash) SelectPossibleRoles(ctx context.Context) ([]core.Role, error) {
	rs := make([]core.Role, 3)
	rs[0] = core.Role{ID: core.RoleAdmin, Name: "Администратор", Const: "RoleAdmin"}
	rs[1] = core.Role{ID: core.RoleAuth, Name: "Авторизованный пользователь", Const: "RoleAuth"}
	rs[2] = core.Role{ID: core.RoleEmployee, Name: "Сотрудник", Const: "RoleEmployee"}

	return rs, nil
}

func (rc *RoleCash) SelectRole(ctx context.Context, r *core.Role) error {
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
