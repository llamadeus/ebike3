package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/llamadeus/ebike3/packages/auth/domain/model"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/in"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/out"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/utils"
)

// AuthService implements the AuthService interface.
type AuthService struct {
	repository out.AuthRepository
}

// Ensure that AuthService implements the AuthService interface.
var _ in.AuthService = (*AuthService)(nil)

// NewAuthService creates a new instance of the AuthService.
func NewAuthService(repository out.AuthRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

// GetAuthBySessionID returns the auth with the given session id.
func (s *AuthService) GetAuthBySessionID(sessionID string) (*model.Auth, error) {
	session, err := s.repository.GetSessionBySessionID(sessionID)
	if err != nil {
		return nil, err
	}

	if session == nil || session.DeletedAt.Valid {
		return nil, micro.NewBadRequestError("invalid session")
	}

	user, err := s.repository.GetUserByID(session.UserID)
	if user == nil {
		return nil, err
	}

	return &model.Auth{
		User:    user,
		Session: session,
	}, nil
}

// RegisterUser registers a new user with the given username, password, and role.
// If the user already exists, an error is returned.
// The password is hashed before being stored in the database.
func (s *AuthService) RegisterUser(username string, password string, role model.UserRole) (*model.Auth, error) {
	existing, err := s.repository.GetUserByUsername(username)
	if existing != nil {
		return nil, micro.NewBadRequestError("user already exists")
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to hash password: %v", err))
	}

	auth, err := s.repository.CreateUserWithSession(username, hashed, role)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create user with session: %v", err))
	}

	return auth, nil
}

// LoginUser logs in the user with the given username and password.
func (s *AuthService) LoginUser(username string, password string) (*model.Auth, error) {
	user, err := s.repository.GetUserByUsername(username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to find user: %v", err))
	}

	if user == nil {
		return nil, micro.NewBadRequestError("invalid username or password")
	}

	if !utils.VerifyPassword(password, user.Password) {
		return nil, micro.NewBadRequestError("invalid username or password")
	}

	session, err := s.repository.CreateSessionAndUpdateLastLogin(user)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create session: %v", err))
	}

	return &model.Auth{
		User:    user,
		Session: session,
	}, nil
}

func (s *AuthService) TerminateSession(sessionID string) error {
	return s.repository.TerminateSession(sessionID)
}
