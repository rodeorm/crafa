package team

func getQueries() map[string]string {
	res := make(map[string]string, 0)

	res["insertTeam"] = `		INSERT INTO cmn.Teams
										  (Name)
										  SELECT $1
										  RETURNING id;`

	res["insertUserTeams"] = `	INSERT INTO cmn.UserTeams
							(UserID, TeamID)
							SELECT $1, $2;`

	res["updateTeam"] = `			UPDATE cmn.Teams
										  SET Name = $2
										  WHERE ID = $1;`

	res["selectTeam"] = `			SELECT id, name
										  FROM cmn.Teams
										  WHERE ID = $1;`

	res["selectAllTeams"] = `			SELECT id, name
											  FROM cmn.Teams;`

	res["selectUserTeams"] = `		SELECT p.id, p.name
											  FROM cmn.Teams AS p
												  INNER JOIN cmn.UserTeams AS up
													  ON p.ID = up.TeamID
											  WHERE up.UserID = $1
											  ;`

	res["deleteTeam"] = `	DELETE FROM cmn.Teams
									  WHERE ID = $1;`

	res["deleteUserTeam"] = `	DELETE FROM cmn.UserTeams
										  WHERE UserID = $1 AND TeamID = $2;`

	res["selectPossibleUserTeams"] = `			SELECT p.id, p.name
													  FROM cmn.Teams AS p
													  LEFT JOIN cmn.UserTeams AS up
													  ON p.ID = up.TeamID AND up.UserID = $1
													  WHERE up.ID IS NULL
													  ;`

	res["selectTeamUsers"] = `			SELECT u.ID AS "user.id", u.Login, u.Name, u.FamilyName, u.PatronName, u.Email
											  FROM cmn.UserTeams AS ut
													  INNER JOIN cmn.Users AS u 
														  ON u.ID  = ut.UserID
											  WHERE ut.TeamID = $1;`
	return res

}
