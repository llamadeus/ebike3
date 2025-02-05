package in

import "github.com/llamadeus/ebike3/packages/auth/domain/model"

type AuthService interface {
	// GetAuthBySessionID returns the auth with the given session id.
	GetAuthBySessionID(sessionID string) (*model.Auth, error)

	// RegisterUser registers a new user and returns it.
	RegisterUser(username string, password string, role model.UserRole) (*model.Auth, error)

	// LoginUser logs in the user with the given username and password.
	LoginUser(username string, password string) (*model.Auth, error)

	// TerminateSession terminates the session with the given id.
	TerminateSession(sessionID string) error
}
