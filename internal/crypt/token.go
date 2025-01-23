package crypt

func GetRoleIDFromTkn(tknStr, jwtKey string) (int, error) {
	cl, err := GetClaims(tknStr, jwtKey)
	if err != nil {
		return 0, err
	}

	return cl.RoleID, nil
}
