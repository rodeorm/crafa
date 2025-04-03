package msg

func getQueries() map[string]string {
	res := make(map[string]string, 0)
	res["insertMsg"] = `INSERT INTO msg.Messages
										 (typeid, categoryid, userid, text, email) 
										  SELECT $1, $2, $3, $4, $5
										  RETURNING id;`

	res["updateMsg"] = `	UPDATE msg.Messages 
																			 SET Used = $2, Queued = $3, SendTime = $4 
																			 WHERE id = $1;`

	res["selectUnsendedMsgs"] = `	SELECT 
																					 id, userid AS "user.id", typeid as "type.id", categoryid as "category.id", text, email 
																					 FROM msg.Messages  
																					 WHERE SendTime IS NULL AND (Queued IS NULL OR Queued = false);`

	res["selectConfMsg"] = `	SELECT 
																				 id, used, queued, sendtime
																				 FROM msg.Messages  
																				 WHERE Used = false AND UserID = $1 AND Text = $2;`

	res["selectAuthMsg"] = `		SELECT 
																				 id, used, queued, sendtime
																				 FROM msg.Messages  
																				 WHERE Used = false AND UserID = $1 AND Text = $2;`

	return res
}
