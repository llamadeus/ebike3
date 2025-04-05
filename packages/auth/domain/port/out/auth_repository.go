package out

import "github.com/llamadeus/ebike3/packages/auth/domain/model"

// AuthRepository is an interface for a repository that handles auth operations.
type AuthRepository interface {
	// CreateUserWithSession creates a new user with the given username, password, and role and logs in the user.
	CreateUserWithSession(username string, password string, role model.UserRole) (*model.Auth, error)

	// CreateSessionAndUpdateLastLogin creates a new session for the user and updates the last login time.
	CreateSessionAndUpdateLastLogin(user *model.User) (*model.Session, error)

	// GetAll returns all registered users.
	GetAll() ([]*model.User, error)

	// GetUserByID returns the user with the given id.
	GetUserByID(id uint64) (*model.User, error)

	// GetUserByUsername returns the user with the given username.
	GetUserByUsername(username string) (*model.User, error)

	// GetSessionByID returns the session with the given id.
	GetSessionByID(id uint64) (*model.Session, error)

	// TerminateSession terminates the session with the given id.
	TerminateSession(id uint64) error
}
