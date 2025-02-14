package postgres

import "github.com/pkg/errors"

func (s *postgresStorage) prepareStmts() error {
	err := s.userPrepareStmts()
	if err != nil {
		return err
	}
	err = s.sessionPrepareStmts()
	if err != nil {
		return err
	}
	err = s.msgPrepareStmts()
	if err != nil {
		return err
	}
	/*
		err = s.epicPrepareStmts()
		if err != nil {
			return err
		}
		err = s.issuePrepareStmts()
		if err != nil {
			return err
		}*/
	err = s.categoryPrepareStmts()
	if err != nil {
		return err
	}
	err = s.projectPrepareStmts()
	if err != nil {
		return err
	}
	/*
		err = s.areaPrepareStmts()
		if err != nil {
			return err
		}

		err = s.iterationPrepareStmts()
		if err != nil {
			return err
		}*/

	return nil
}

func (s *postgresStorage) sessionPrepareStmts() error {
	insertSession, err := s.DB.Preparex(`INSERT INTO cmn.Sessions
	(userid, logintime, actiontime) 
	 SELECT $1, $2, $3
	 RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertSession")
	}

	s.preparedStatements["insertSession"] = insertSession
	return nil
}

func (s *postgresStorage) userPrepareStmts() error {
	insertUser, err := s.DB.Preparex(`	INSERT INTO cmn.Users 
	(roleid, login, password, name, familyname, patronname, email, phone) 
	SELECT $1, $2, $3, $4, $5, $6, $7, $8
	RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertUser")
	}

	selectUser, err := s.DB.Preparex(`SELECT 
									  roleid AS "role.id", login, name, familyname, patronname, email, phone
									  FROM cmn.Users
									  WHERE id = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectUser")
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
	s.preparedStatements["insertUser"] = insertUser
	s.preparedStatements["selectUser"] = selectUser
	s.preparedStatements["updateUserRole"] = updateUserRole
	s.preparedStatements["baseAuthUser"] = baseAuthUser
	s.preparedStatements["selectAllUsers"] = selectAllUsers
	s.preparedStatements["updateUser"] = updateUser
	s.preparedStatements["changeUserPassword"] = changeUserPassword
	return nil
}

func (s *postgresStorage) msgPrepareStmts() error {
	insertMsg, err := s.DB.Preparex(`INSERT INTO msg.Messages
	(typeid, categoryid, userid, text, email) 
	 SELECT $1, $2, $3, $4, $5
	 RETURNING id;`)
	if err != nil {
		errors.Wrap(err, "insertMsg")
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

	selectAuthMsg, err := s.DB.Preparex(`		SELECT 
											id, used, queued, sendtime
											FROM msg.Messages  
											WHERE Used = false AND UserID = $1 AND Text = $2;`)
	if err != nil {
		return errors.Wrap(err, "selectAuthMsg")
	}

	s.preparedStatements["insertMsg"] = insertMsg
	s.preparedStatements["updateMsg"] = updateMsg
	s.preparedStatements["selectUnsendedMsgs"] = selectUnsendedMsgs
	s.preparedStatements["selectConfMsg"] = selectConfMsg
	s.preparedStatements["selectAuthMsg"] = selectAuthMsg
	return nil
}

/*
	func (s *postgresStorage) issuePrepareStmts() error {
		insertIssue, err := s.DB.Preparex(`	INSERT INTO data.Issues
											()
											SELECT $1, $2, $3, $4, $5, $6, $7, $8;`)
		if err != nil {
			return errors.Wrap(err, "insertUser")
		}
		/*
			InsertIssue(context.Context, *Issue) error
			SelectIssue(context.Context, *User) (*Issue, error)
			UpdateIssue(context.Context, *Issue) error
			DeleteIssue(context.Context, *Issue) error
			InsertIssueComment(context.Context, *Issue, *Comment) error
			DeleteIssueComment(context.Context, *Issue, *Comment) error
			UpdateIssueComment(context.Context, *Issue, *Comment) error
			SelectAllIssueComments(context.Context, *Issue) error
			SelectAllProjectIssues(context.Context, *Project) ([]Issue, error)
		s.preparedStatements["insertIssue"] = insertIssue
		return nil
	}
*/
func (s *postgresStorage) projectPrepareStmts() error {

	insertProject, err := s.DB.Preparex(`		INSERT INTO data.Projects
												(Name)
												SELECT $1
												RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertProject")
	}

	insertUserProject, err := s.DB.Preparex(`		INSERT INTO data.UserProjects
	(UserID, ProjectID)
	SELECT $1, $2;`)
	if err != nil {
		return errors.Wrap(err, "insertUserProject")
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

	selectUserProjects, err := s.DB.Preparex(`		SELECT p.id, p.name
													FROM data.Projects AS p
														INNER JOIN data.UserProjects AS up
															ON p.ID = up.ProjectID
													WHERE up.UserID = $1
													;`)
	if err != nil {
		return errors.Wrap(err, "selectUserProjects")
	}

	deleteProject, err := s.DB.Preparex(`	DELETE FROM data.Projects
											WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "deleteProject")
	}

	deleteUserProject, err := s.DB.Preparex(`	DELETE FROM data.UserProjects
												WHERE UserID = $1 AND ProjectID = $2;`)
	if err != nil {
		return errors.Wrap(err, "deleteUserProject")
	}

	selectPossibleUserProjects, err := s.DB.Preparex(`		SELECT p.id, p.name
															FROM data.Projects AS p
															LEFT JOIN data.UserProjects AS up
															ON p.ID = up.ProjectID AND up.UserID = $1
															WHERE up.ID IS NULL
															;`)
	if err != nil {
		return errors.Wrap(err, "selectPossibleUserProjects")
	}

	s.preparedStatements["insertProject"] = insertProject
	s.preparedStatements["updateProject"] = updateProject
	s.preparedStatements["selectProject"] = selectProject
	s.preparedStatements["selectAllProjects"] = selectAllProjects
	s.preparedStatements["deleteProject"] = deleteProject
	s.preparedStatements["selectUserProjects"] = selectUserProjects
	s.preparedStatements["deleteUserProject"] = deleteUserProject
	s.preparedStatements["selectPossibleUserProjects"] = selectPossibleUserProjects
	s.preparedStatements["insertUserProject"] = insertUserProject

	return nil
}

/*
func (s *postgresStorage) areaPrepareStmts() error {
	return nil
}

func (s *postgresStorage) iterationPrepareStmts() error {
	return nil
}
*/
