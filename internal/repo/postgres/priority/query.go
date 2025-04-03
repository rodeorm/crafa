package priority

func getQueries() map[string]string {
	res := make(map[string]string, 0)
	res["insertPriority"] = `		INSERT INTO ref.Priorities
										  (Name, LevelID)
										  SELECT $1, $2
										  RETURNING id;`

	res["updatePriority"] = `		UPDATE ref.Priorities
										  SET Name = $2, LevelID = $3
										  WHERE ID = $1;`

	res["selectPriority"] = `		SELECT c.id, c.name, l.ID AS "level.id", l.name AS "level.name"
										  FROM ref.Priorities AS c
											  INNER JOIN ref.Levels AS l ON l.ID = c.LevelID
										  WHERE c.ID = $1;`

	res["selectAllPriorities"] = `		SELECT c.id as "id", c.name as "name", l.id AS "level.id", l.name AS "level.name"
											  FROM ref.Priorities AS c 
												  INNER JOIN ref.Levels AS l ON l.ID = c.LevelID ;`
	res["selectLevelPriorities"] = `	SELECT c.id as "id", c.name as "name", l.id AS "level.id", l.name AS "level.name"
											  FROM ref.Priorities AS c 
												  INNER JOIN ref.Levels AS l ON l.ID = c.LevelID 
											  WHERE levelid = $1;`

	res["deletePriority"] = `	DELETE FROM ref.Priorities
									  WHERE ID = $1;`
	return res
}
