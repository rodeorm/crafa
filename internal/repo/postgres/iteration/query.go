package iteration

func getQueries() map[string]string {
	res := make(map[string]string, 0)

	res["updateIteration"] = `		UPDATE ref.Iterations
										 SET Name = $2, LevelID = $3, ParentID = $4, Year = $5, Month = $6
										 WHERE ID = $1;`

	res["insertIteration"] = `		INSERT INTO ref.Iterations
										 (Name, LevelID, ParentID, Year, Month)
										 SELECT $1, $2, $3, $4, $5
										 RETURNING id;`

	res["selectIteration"] = `		SELECT i.id, i.name, i.year, i.month,
												 l.id AS "level.id", l.name AS "level.name", 
												 COALESCE(p.id,0) AS "parent.id", COALESCE(p.name,'-') AS "parent.name"
										 FROM ref.Iterations AS i
											 LEFT JOIN ref.Iterations AS p ON p.ID = i.ParentID
											 INNER JOIN ref.Levels AS l ON l.ID = i.levelID
										 WHERE i.ID = $1;`

	res["selectAllIterations"] = `			SELECT  i.id, i.name, i.year, i.month,
													 l.id AS "level.id", l.name AS "level.name", 
													 COALESCE(p.id,0) AS "parent.id", COALESCE(p.name,'-') AS "parent.name"
												 FROM ref.Iterations AS i
													 LEFT JOIN ref.Iterations AS p ON p.ID = i.ParentID
												 INNER JOIN ref.Levels AS l ON l.ID = i.levelID;`

	res["deleteIteration"] = `	DELETE FROM ref.Iterations
									 WHERE ID = $1;`

	res["selectPossibleLevelIterations"] = `			SELECT i.id, i.name, i.year, i.month,
																 l.id AS "level.id", l.name AS "level.name", 
																 COALESCE(p.id,0) AS "parent.id", COALESCE(p.name,'-') AS "parent.name"
															 FROM ref.Iterations AS i
																 LEFT JOIN ref.Iterations AS p ON p.ID = i.ParentID
																 INNER JOIN ref.Levels AS l ON l.ID = i.levelID
															 WHERE i.LevelID = $1;`

	res["selectPossibleParentIterations"] = `			SELECT 	p.id AS "id", p.name AS "name",
																	 l.id AS "level.id", l.name AS "level.name"
															 FROM ref.Iterations AS i
																 INNER JOIN ref.Iterations AS p ON p.LevelID = i.LevelID - 1
																 INNER JOIN ref.Levels AS l ON l.ID = p.LevelID
															 WHERE i.ID = $1;`

	return res
}
