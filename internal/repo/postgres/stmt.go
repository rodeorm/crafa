package postgres

func (s *postgresStorage) prepareStatements() error {

	insertUser, err := s.DB.Preparex(`	INSERT INTO cmn.Users 
										(roleid, login, password, name, familyname, patronname, email, phone) 
										SELECT $1, $2, $3, $4, $5, $6, $7, $8
										RETURNING id;`)
	if err != nil {
		return err
	}

	insertEmail, err := s.DB.Preparex(`INSERT INTO msg.Emails
									   (typeid, userid, text, email) 
										SELECT $1, $2, $3, $4
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

	/*
		stmtUpdateUser, err := s.DB.Preparex(`UPDATE cmn.Users SET name = $2, email = $3, phone = $4, password = $5, verified = $6 WHERE ID = $1;`)
		if err != nil {
			return err
		}
		stmtAuthUser, err := s.DB.Preparex(`SELECT Login, Password, ID, Email, RoleID AS role.ID FROM cmn.Users WHERE Login = $1;`)
		if err != nil {
			return err
		}
		stmtVerifyUser, err := s.DB.Preparex(`SELECT e.id FROM cmn.emails AS e INNER JOIN cmn.Users AS u ON u.id = e.UserID WHERE u.Login = $1 AND e.sendeddate + ($2 * INTERVAL '1 hour') > NOW()
		AND e.Used = false AND e.OTP = $3; `)
		if err != nil {
			return err
		}

		stmtStartSession, err := s.DB.Preparex(`INSERT INTO cmn.Sessions (UserID, LoginDate, LastActionDate) SELECT $1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP`)
		if err != nil {
			return err
		}
		stmtUpdateSession, err := s.DB.Preparex(`UPDATE cmn.Sessions SET LastActionDate = CURRENT_TIMESTAMP WHERE id = $1;`)
		if err != nil {
			return err
		}
		stmtEndSession, err := s.DB.Preparex(`UPDATE cmn.Sessions SET LogoutDate = CURRENT_TIMESTAMP WHERE id = $1;`)
		if err != nil {
			return err
		}
		stmtAddEmail, err := s.DB.Preparex(`INSERT INTO cmn.Emails (UserID, OTP, Email) SELECT $1, $2, $3 RETURNING id;`)
		if err != nil {
			return err
		}
		stmtUpdateEmail, err := s.DB.Preparex(`UPDATE cmn.Emails SET OTP = $2, Email = $3, SendedDate = $4, Used = $5, Queued = $6 WHERE id = $1;`)
		if err != nil {
			return err
		}
		stmpSelectEmailForSending, err := s.DB.Preparex(`SELECT id, userid, otp, email AS destination FROM cmn.Emails WHERE SendedDate IS NULL AND (Queued IS NULL OR Queued = false);`)
		if err != nil {
			return err
		}
	*/
	s.preparedStatements["insertUser"] = insertUser
	s.preparedStatements["insertEmail"] = insertEmail
	s.preparedStatements["insertSession"] = insertSession
	s.preparedStatements["selectUser"] = selectUser
	/*
		s.preparedStatements["UpdateUser"] = stmtUpdateUser
		s.preparedStatements["AuthUser"] = stmtAuthUser
		s.preparedStatements["VerifyUser"] = stmtVerifyUser
		s.preparedStatements["StartSession"] = stmtStartSession
		s.preparedStatements["UpdateSession"] = stmtUpdateSession
		s.preparedStatements["EndSession"] = stmtEndSession
		s.preparedStatements["AddEmail"] = stmtAddEmail
		s.preparedStatements["UpdateEmail"] = stmtUpdateEmail
		s.preparedStatements["SelectEmailForSending"] = stmpSelectEmailForSending
	*/
	return nil
}
