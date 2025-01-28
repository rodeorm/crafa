package postgres

func (s *postgresStorage) prepareStatements() error {

	insertUser, err := s.DB.Preparex(`	INSERT INTO cmn.Users 
										(roleid, login, password, name, familyname, patronname, email, phone) 
										SELECT $1, $2, $3, $4, $5, $6, $7, $8
										RETURNING id;`)
	if err != nil {

		return err
	}

	insertMsg, err := s.DB.Preparex(`INSERT INTO msg.Messages
									   (typeid, categoryid, userid, text, email) 
										SELECT $1, $2, $3, $4, $5
										RETURNING id;`)
	if err != nil {
		return err
	}

	insertSession, err := s.DB.Preparex(`INSERT INTO cmn.Sessions
									   (userid, logintime, actiontime) 
										SELECT $1, $2, $3
										RETURNING id;`)
	if err != nil {
		return err
	}

	selectUser, err := s.DB.Preparex(`SELECT 
									  roleid AS "role.id", name, familyname, patronname, email, phone
									  FROM cmn.Users
									  WHERE id = $1`)
	if err != nil {
		return err
	}

	updateMsg, err := s.DB.Preparex(`	UPDATE msg.Messages 
										SET Used = $2, Queued = $3, SendTime = $4 
										WHERE id = $1;`)
	if err != nil {
		return err
	}
	selectUnsendedMsgs, err := s.DB.Preparex(`	SELECT 
												id, userid AS "user.id", typeid as "type.id", categoryid as "category.id", text, email 
											    FROM msg.Messages  
												WHERE SendTime IS NULL AND (Queued IS NULL OR Queued = false);`)
	if err != nil {
		return err
	}
	selectConfMsg, err := s.DB.Preparex(`	SELECT 
											id, used, queued, sendtime
											FROM msg.Messages  
											WHERE Used = false AND UserID = $1 AND Text = $2;`)
	if err != nil {
		return err
	}
	updateUserRole, err := s.DB.Preparex(`	UPDATE cmn.Users
											SET  RoleID = $2
											WHERE ID = $1;`)
	if err != nil {
		return err
	}

	s.preparedStatements["insertUser"] = insertUser
	s.preparedStatements["insertMsg"] = insertMsg
	s.preparedStatements["insertSession"] = insertSession
	s.preparedStatements["selectUser"] = selectUser
	s.preparedStatements["updateMsg"] = updateMsg
	s.preparedStatements["selectUnsendedMsgs"] = selectUnsendedMsgs
	s.preparedStatements["selectConfMsg"] = selectConfMsg
	s.preparedStatements["updateUserRole"] = updateUserRole

	return nil
}
