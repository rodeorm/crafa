package epic

func getQueries() map[string]string {
	res := make(map[string]string, 0)
	res["insertPriority"] = `		INSERT INTO ref.Priorities
										  (Name, LevelID)
										  SELECT $1, $2
										  RETURNING id;`
	res["insertProject"] = `		INSERT INTO data.Projects
										  (Name)
										  SELECT $1
										  RETURNING id;`
	res["insertUserProject"] = `		INSERT INTO data.UserProjects
									(UserID, ProjectID)
									SELECT $1, $2;`

	res["updateProject"] = `		UPDATE data.Projects
										  SET Name = $2
										  WHERE ID = $1;`

	res["selectProject"] = `		SELECT id, name
										  FROM data.Projects
										  WHERE ID = $1;`
	res["selectProjectEpics"] = `		SELECT id, name
										  FROM data.Epics AS e
										  WHERE ProjectID = $1;`

	res["selectAllProjects"] = `		SELECT id, name
											  FROM data.Projects;`
	res["selectUserProjects"] = `		SELECT p.id, p.name
											  FROM data.Projects AS p
												  INNER JOIN data.UserProjects AS up
													  ON p.ID = up.ProjectID
											  WHERE up.UserID = $1
											  ;`
	res["deleteProject"] = `	DELETE FROM data.Projects
									  WHERE ID = $1;`
	res["deleteUserProject"] = `	DELETE FROM data.UserProjects
										  WHERE UserID = $1 AND ProjectID = $2;`
	res["selectPossibleUserProjects"] = `		SELECT p.id, p.name
													  FROM data.Projects AS p
													  LEFT JOIN data.UserProjects AS up
													  ON p.ID = up.ProjectID AND up.UserID = $1
													  WHERE up.ID IS NULL
													  ;`
	return res
}
