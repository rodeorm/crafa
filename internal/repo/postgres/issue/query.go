package issue

func getQueries() map[string]string {
	res := make(map[string]string, 0)
	res["insertIssue"] = `		INSERT INTO ref.Priorities
										  (Name, LevelID)
										  SELECT $1, $2
										  RETURNING id;`
	res["insertIssue"] = `		INSERT INTO data.Issues
										  (Name)
										  SELECT $1
										  RETURNING id;`
	res["insertUserIssue"] = `		INSERT INTO data.UserIssues
									(UserID, IssueID)
									SELECT $1, $2;`

	res["updateIssue"] = `		UPDATE data.Issues
										  SET Name = $2
										  WHERE ID = $1;`

	res["selectIssue"] = `		SELECT id, name
										  FROM data.Issues
										  WHERE ID = $1;`

	res["selectEpicIssuess"] = `		SELECT id, name
										  FROM data.Issues AS i
										  WHERE i.EpicID = $1;`

	res["selectAllIssues"] = `		SELECT id, name
											  FROM data.Issues;`
	res["selectUserIssues"] = `		SELECT p.id, p.name
											  FROM data.Issues AS p
												  INNER JOIN data.UserIssues AS up
													  ON p.ID = up.IssueID
											  WHERE up.UserID = $1
											  ;`
	res["deleteIssue"] = `	DELETE FROM data.Issues
									  WHERE ID = $1;`
	res["deleteUserIssue"] = `	DELETE FROM data.UserIssues
										  WHERE UserID = $1 AND IssueID = $2;`
	res["selectUserIssues"] = `		SELECT p.id, p.name
													  FROM data.Issues AS p
													  	LEFT JOIN data.UserIssues AS up
													  ON p.ID = up.IssueID AND up.UserID = $1
													  WHERE up.ID IS NULL
													  ;`
	return res
}
