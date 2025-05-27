package category

func getQueries() map[string]string {
	res := make(map[string]string, 0)
	res["insertCategory"] = `		INSERT INTO ref.Categories
										 (Name, LevelID)
										 SELECT $1, $2
										 RETURNING id;`

	res["updateCategory"] = `		UPDATE ref.Categories
										 SET Name = $2, LevelID = $3
										 WHERE ID = $1;`

	res["selectCategory"] = `		SELECT c.id, c.name, l.ID AS "level.id", l.name AS "level.name"
										 FROM ref.Categories AS c
											 INNER JOIN ref.Levels AS l ON l.ID = c.LevelID
										 WHERE c.ID = $1;`

	res["selectAllCategories"] = `		SELECT c.id as "id", c.name as "name", l.id AS "level.id", l.name AS "level.name"
											 FROM ref.Categories AS c 
												 INNER JOIN ref.Levels AS l ON l.ID = c.LevelID ;`

	res["selectLevelCategories"] = `	SELECT c.id as "id", c.name as "name", l.id AS "level.id", l.name AS "level.name"
											 FROM ref.Categories AS c 
												 INNER JOIN ref.Levels AS l ON l.ID = c.LevelID 
											 WHERE levelid = $1;`

	res["deleteCategory"] = `	DELETE FROM ref.Categories
									 WHERE ID = $1;`

	return res
}
