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
									  roleid AS "role.id", login, name, familyname, patronname, email, phone
									  FROM cmn.Users
									  WHERE id = $1;`)
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

	selectAuthMsg, err := s.DB.Preparex(`		SELECT 
											id, used, queued, sendtime
											FROM msg.Messages  
											WHERE Used = false AND UserID = $1 AND Text = $2;`)
	if err != nil {
		return errors.Wrap(err, "selectAuthMsg")
	}

	selectAllUsers, err := s.DB.Preparex(`		SELECT 
												u.id AS "user.id", roleid AS "role.id", r.Name AS "role.name", login, u.name, familyname, patronname, email, phone
												FROM cmn.Users AS u
													INNER JOIN cmn.Roles AS r ON r.ID = u.RoleID;`)
	if err != nil {
		return errors.Wrap(err, "selectAllUsers")
	}

	updateUser, err := s.DB.Preparex(`		UPDATE cmn.Users 
											SET roleid = $2, login = $3, name = $4, familyname = $5, patronname = $6, email = $7, phone = $8
											WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "updateUser")
	}
	changeUserPassword, err := s.DB.Preparex(`	UPDATE cmn.Users 
												SET password = $2
												WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "changeUserPassword")
	}

	insertProject, err := s.DB.Preparex(`		INSERT INTO data.Projects
												(Name)
												SELECT $1
												RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertProject")
	}

	updateProject, err := s.DB.Preparex(`		UPDATE data.Projects 
												SET Name = $2
												WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "updateProject")
	}

	selectProject, err := s.DB.Preparex(`		SELECT id, name
												FROM data.Projects
												WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectProject")
	}

	selectAllProjects, err := s.DB.Preparex(`		SELECT id, name 
													FROM data.Projects;`)
	if err != nil {
		return errors.Wrap(err, "selectAllProjects")
	}

	deleteProject, err := s.DB.Preparex(`	DELETE FROM data.Projects
											WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "deleteProject")
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
	s.preparedStatements["selectAuthMsg"] = selectAuthMsg
	s.preparedStatements["selectAllUsers"] = selectAllUsers
	s.preparedStatements["updateUser"] = updateUser
	s.preparedStatements["changeUserPassword"] = changeUserPassword
	s.preparedStatements["insertProject"] = insertProject
	s.preparedStatements["updateProject"] = updateProject
	s.preparedStatements["selectProject"] = selectProject
	s.preparedStatements["selectAllProjects"] = selectAllProjects
	s.preparedStatements["deleteProject"] = deleteProject

	return nil
}
