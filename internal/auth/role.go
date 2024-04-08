package auth

import (
	"context"

	"money/internal/core"
)

// SelectRoleData возвращает данные роли из БД. В случае отсутствия данных о роли в БД вернется ErrNoRows
func (s *postgresStorage) SelectRoleData(ctx context.Context, r *core.Role) error {
	err := s.DB.QueryRowContext(ctx, "SELECT C_Name FROM cmn.Roles WHERE ID = $1", r.ID).Scan(&r.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *postgresStorage) SelectAllRoles(ctx context.Context) (*[]core.Role, error) {
	rows, err := s.DB.QueryContext(ctx, "SELECT id, C_Name FROM cmn.roles")
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()
	roles := make([]core.Role, 0, 1)
	for rows.Next() {
		var (
			role core.Role
		)
		err = rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, err
		}
		err := s.SelectRoleData(ctx, &role)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return &roles, nil
}
