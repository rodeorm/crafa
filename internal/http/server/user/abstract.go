package user

import (
	"context"
	"money/internal/core"
	"net/http"
	"time"
)

type SessionManager interface {
	GetSession(r *http.Request) (*core.Session, error)
}

type UserStorager interface {
	RegUser(ctx context.Context, u *core.User, domain string) (*core.Session, error)       //	RegUser добавляет нового пользователя. Возвращает письмо для подтверждения адреса электронной почты и сессию
	SelectUser(ctx context.Context, u *core.User) error                                    //	SelectUser возвращает данные пользователя
	ConfirmUserEmail(ctx context.Context, userID int, otp string) error                    //	ConfirmUserEmail подтверждает адрес электронной почты для нового пользователя. Возвращает ошибку, если подтверждение не удалось
	BaseAuthUser(context.Context, *core.User) error                                        //	BaseAuthUser авторизует пользователя через базовую аутентификацию по логину-паролю
	AdvAuthUser(context.Context, *core.User, string, time.Duration) (*core.Session, error) //	AdvAuthUser авторизует пользователя, прошедшего базовую аутентификацию по одноразовому паролю
	UpdateUser(context.Context, *core.User) error                                          //	UpdateUser обновляет данные пользователя
	SelectAllUsers(ctx context.Context) ([]core.User, error)                               //   SelectAllUsers возвращает данные всех пользователей
}

type CookieManager interface {
	NewCookieWithSession(s *core.Session) (*http.Cookie, error)
}

type RoleStorager interface {
	SelectPossibleRoles(context.Context) ([]core.Role, error)
	SelectRole(context.Context, *core.Role) error
}

type ProjectStorager interface { //TODO: разбить на два интерфейса
	InsertProject(context.Context, *core.Project) error
	InsertUserProject(ctx context.Context, userID, projectID int) error
	UpdateProject(context.Context, *core.Project) error
	SelectProject(context.Context, *core.Project) error
	SelectAllProjects(context.Context) ([]core.Project, error)
	SelectUserProjects(context.Context, *core.User) ([]core.Project, error)
	DeleteProject(context.Context, *core.Project) error
	DeleteUserProject(context.Context, *core.User, *core.Project) error
	SelectPossibleNewUserProjects(context.Context, *core.User) ([]core.Project, error)
	SelectAllProjectEpics(context.Context, *core.Project) ([]core.Epic, error)
}

type TeamStorager interface {
	InsertTeam(ctx context.Context, p *core.Team) error
	SelectTeam(ctx context.Context, p *core.Team) error
	UpdateTeam(ctx context.Context, p *core.Team) error
	SelectAllTeams(ctx context.Context) ([]core.Team, error)
	SelectUserTeams(ctx context.Context, u *core.User) ([]core.Team, error)
	DeleteTeam(ctx context.Context, p *core.Team) error
	DeleteUserTeam(ctx context.Context, u *core.User, p *core.Team) error
	InsertUserTeams(ctx context.Context, userID, TeamID int) error
	SelectPossibleNewUserTeams(ctx context.Context, u *core.User) ([]core.Team, error)
	SelectAllTeamEpics(ctx context.Context, c *core.Team) ([]core.Epic, error)
}
