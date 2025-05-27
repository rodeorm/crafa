package area

func getQueries() map[string]string {
	res := make(map[string]string, 0)

	res["InsertArea"] = `	INSERT INTO ref.Areas
										(name, levelid) 
	 									SELECT $1, $2
	 									RETURNING id;`

	res["UpdateArea"] = `	UPDATE ref.Areas
										SET name = $2, levelid = $3 
	 									WHERE ID = $1;`

	res["DeleteArea"] = `	DELETE FROM ref.Areas 
	 									WHERE ID = $1;`

	res["SelectArea"] = `	SELECT  
	 									a.Name, l.id AS "level.id", l.name AS "level.name" 
										FROM ref.Areas AS a
												INNER JOIN ref.Levels AS l ON l.ID = a.LevelID
	 									WHERE a.ID = $1;`

	res["SelectAllAreas"] = `	SELECT  
	 										a.ID, a.Name, l.ID AS "level.id", l.Name AS "level.name"
											FROM ref.Areas AS a
												INNER JOIN ref.Levels AS l ON l.ID = a.LevelID;`

	res["SelectAllLevelAreas"] = `	SELECT  
												ID, Name
										  		FROM ref.Areas
												WHERE LevelID = $1;`

	return res
}
