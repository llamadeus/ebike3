package out

import "github.com/llamadeus/ebike3/packages/auth/domain/model"

// AuthRepository is an interface for a repository that handles auth operations.
type AuthRepository interface {
	CreateUserWithSession(username string, password string, role model.UserRole) (*model.Auth, error)

	CreateSessionAndUpdateLastLogin(user *model.User) (*model.Session, error)

	GetUserByID(id uint64) (*model.User, error)

	GetUserByUsername(username string) (*model.User, error)

	GetSessionByID(id uint64) (*model.Session, error)

	GetSessionBySessionID(sessionID string) (*model.Session, error)

	TerminateSession(sessionID string) error
}
