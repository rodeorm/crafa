package user

func getQueries() map[string]string {
	res := make(map[string]string, 0)

	res["insertTeam"] = `		INSERT INTO cmn.Teams
										  (Name)
										  SELECT $1
										  RETURNING id;`
	res["insertMsg"] = `INSERT INTO msg.Messages
										  (typeid, categoryid, userid, text, email) 
										   SELECT $1, $2, $3, $4, $5
										   RETURNING id;`
	res["updateMsg"] = `	UPDATE msg.Messages 
										   SET Used = $2, Queued = $3, SendTime = $4 
										   WHERE id = $1;`

	res["selectConfMsg"] = `	SELECT 
										   id, used, queued, sendtime
										   FROM msg.Messages  
										   WHERE Used = false AND UserID = $1 AND Text = $2;`

	res["selectAuthMsg"] = `		SELECT 
										   id, used, queued, sendtime
										   FROM msg.Messages  
										   WHERE Used = false AND UserID = $1 AND Text = $2;`
	res["insertSession"] = `INSERT INTO cmn.Sessions
							(userid, logintime, actiontime) 
	 						SELECT $1, $2, $3
	 						RETURNING id;`

	res["insertUser"] = `	INSERT INTO cmn.Users 
							(roleid, login, password, name, familyname, patronname, email, phone) 
							SELECT $1, $2, $3, $4, $5, $6, $7, $8
	 						RETURNING id;`

	res["selectUser"] = `SELECT 
   						roleid AS "role.id", login, name, familyname, patronname, email, phone
   						FROM cmn.Users
   						WHERE id = $1;`

	res["updateUserRole"] = `	UPDATE cmn.Users
								SET  RoleID = $2
								WHERE ID = $1;`

	res["baseAuthUser"] = `
		 						SELECT id as "user.id", password, roleid AS "role.id", name, familyname, patronname, email, phone
		 						FROM cmn.Users
		 						WHERE login = $1;`

	res["selectAllUsers"] = `		SELECT 
			 						u.id AS "user.id", roleid AS "role.id", r.Name AS "role.name", login, u.name, familyname, patronname, email, phone
									 FROM cmn.Users AS u
									INNER JOIN cmn.Roles AS r ON r.ID = u.RoleID;`

	res["updateUser"] = `			UPDATE cmn.Users 
		 							SET roleid = $2, login = $3, name = $4, familyname = $5, patronname = $6, email = $7, phone = $8
		 							WHERE ID = $1;`

	res["changeUserPassword"] = `	UPDATE cmn.Users 
			 						SET password = $2
			 						WHERE ID = $1;`
	return res
}
