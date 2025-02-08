package postgres

import "github.com/pkg/errors"

func (s *postgresStorage) prepareStatements() error {

	insertUser, err := s.DB.Preparex(`	INSERT INTO cmn.Users 
										(roleid, login, password, name, familyname, patronname, email, phone) 
										SELECT $1, $2, $3, $4, $5, $6, $7, $8
										RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertUser")
	}

	insertMsg, err := s.DB.Preparex(`INSERT INTO msg.Messages
									   (typeid, categoryid, userid, text, email) 
										SELECT $1, $2, $3, $4, $5
										RETURNING id;`)
	if err != nil {
		errors.Wrap(err, "insertMsg")
	}

	insertSession, err := s.DB.Preparex(`INSERT INTO cmn.Sessions
									   (userid, logintime, actiontime) 
										SELECT $1, $2, $3
										RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertSession")
	}

	selectUser, err := s.DB.Preparex(`SELECT 
									  roleid AS "role.id", name, familyname, patronname, email, phone
									  FROM cmn.Users
									  WHERE id = $1`)
	if err != nil {
		return errors.Wrap(err, "selectUser")
	}

	updateMsg, err := s.DB.Preparex(`	UPDATE msg.Messages 
										SET Used = $2, Queued = $3, SendTime = $4 
										WHERE id = $1;`)
	if err != nil {
		return errors.Wrap(err, "updateMsg")
	}
	selectUnsendedMsgs, err := s.DB.Preparex(`	SELECT 
												id, userid AS "user.id", typeid as "type.id", categoryid as "category.id", text, email 
											    FROM msg.Messages  
												WHERE SendTime IS NULL AND (Queued IS NULL OR Queued = false);`)
	if err != nil {
		return errors.Wrap(err, "selectUnsendedMsgs")
	}
	selectConfMsg, err := s.DB.Preparex(`	SELECT 
											id, used, queued, sendtime
											FROM msg.Messages  
											WHERE Used = false AND UserID = $1 AND Text = $2;`)
	if err != nil {
		return errors.Wrap(err, "selectConfMsg")
	}
	updateUserRole, err := s.DB.Preparex(`	UPDATE cmn.Users
											SET  RoleID = $2
											WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "updateUserRole")
	}

	baseAuthUser, err := s.DB.Preparex(`
											SELECT id as "user.id", password, roleid AS "role.id", name, familyname, patronname, email, phone
											FROM cmn.Users
											WHERE login = $1;`)
	if err != nil {
		return errors.Wrap(err, "baseAuthUser")
	}

	advAuthUser, err := s.DB.Preparex(`		SELECT 
											id, used, queued, sendtime
											FROM msg.Messages  
											WHERE Used = false AND UserID = $1 AND Text = $2;`)
	if err != nil {
		return errors.Wrap(err, "advAuthUser")
	}

	s.preparedStatements["insertUser"] = insertUser
	s.preparedStatements["insertMsg"] = insertMsg
	s.preparedStatements["insertSession"] = insertSession
	s.preparedStatements["selectUser"] = selectUser
	s.preparedStatements["updateMsg"] = updateMsg
	s.preparedStatements["selectUnsendedMsgs"] = selectUnsendedMsgs
	s.preparedStatements["selectConfMsg"] = selectConfMsg
	s.preparedStatements["updateUserRole"] = updateUserRole
	s.preparedStatements["baseAuthUser"] = baseAuthUser
	s.preparedStatements["advAuthUser"] = advAuthUser

	return nil
}
