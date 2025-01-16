package crypt

func GetRoleIDFromTkn(tknStr string) (int, error) {
	cl, err := GetClaims(tknStr)
	if err != nil {
		return 0, err
	}

	return cl.RoleID, nil
}
