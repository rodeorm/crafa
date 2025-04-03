package status

func getQueries() map[string]string {
	res := make(map[string]string, 0)
	res["insertPriority"] = `		INSERT INTO ref.Priorities
										  (Name, LevelID)
										  SELECT $1, $2
										  RETURNING id;`
	res["insertStatus"] = `	INSERT INTO ref.Statuses
										  (name, levelid) 
										   SELECT $1, $2
										   RETURNING id;`

	res["insertStatusHierarchy"] = `	INSERT INTO ref.StatusHierarchy
													  (parent, child) 
													   SELECT $1, $2;`

	res["updateStatus"] = `	UPDATE ref.Statuses
										  SET name = $2, levelid = $3 
										   WHERE ID = $1;`

	res["deleteStatus"] = `	DELETE FROM ref.Statuses 
										   WHERE ID = $1;`

	res["selectStatus"] = `	SELECT  
										   a.Name, l.id AS "level.id", l.name AS "level.name" 
										  FROM ref.Statuses AS a
												  INNER JOIN ref.Levels AS l ON l.ID = a.LevelID
										   WHERE a.ID = $1;`

	res["selectAllStatuses"] = `	SELECT  
											   a.ID, a.Name, l.ID AS "level.id", l.Name AS "level.name"
											  FROM ref.Statuses AS a
												  INNER JOIN ref.Levels AS l ON l.ID = a.LevelID;`

	res["selectAllLevelStatuses"] = `	SELECT  
												  ID, Name
													FROM ref.Statuses
												  WHERE LevelID = $1;`

	res["selectFirstLevelStatuses"] = `	SELECT  
															  s.ID, s.Name
															FROM ref.Statuses AS s
															  LEFT JOIN ref.StatusHierarchy AS sh ON sh.Child = s.ID
														  WHERE LevelID = $1 AND sh.Parent IS NULL;`

	res["selectParents"] = `	SELECT  
														  s.ID, s.Name
														FROM ref.Statuses AS s
														  INNER JOIN ref.StatusHierarchy AS sh ON sh.Child = s.ID
													  WHERE LevelID = $1;`

	res["selectPossibleParents"] = `	SELECT  
														  s.ID, s.Name
														FROM ref.Statuses AS s
														  LEFT JOIN ref.StatusHierarchy AS sh ON sh.Child = s.ID
													  WHERE LevelID = $1 AND sh.Parent IS NULL;`

	res["selectChildren"] = `	SELECT  
													  s.ID, s.Name
													FROM ref.Statuses AS s
													  INNER JOIN ref.StatusHierarchy AS sh ON sh.Child = s.ID
												  WHERE LevelID = $1;`

	res["selectPossibleChildren"] = `	SELECT  
													  s.ID, s.Name
													FROM ref.Statuses AS s
													  LEFT JOIN ref.StatusHierarchy AS sh ON sh.Child = s.ID
												  WHERE LevelID = $1 AND sh.Parent IS NULL;`

	return res

}
