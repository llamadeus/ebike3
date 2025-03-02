package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/llamadeus/ebike3/packages/auth/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/auth/domain/events"
	"github.com/llamadeus/ebike3/packages/auth/domain/model"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/in"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/out"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

// AuthService implements the AuthService interface.
type AuthService struct {
	kafka      micro.Kafka
	repository out.AuthRepository
}

type UserRegisteredEvent struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// Ensure that AuthService implements the AuthService interface.
var _ in.AuthService = (*AuthService)(nil)

// NewAuthService creates a new instance of the AuthService.
func NewAuthService(kafka micro.Kafka, repository out.AuthRepository) *AuthService {
	return &AuthService{
		kafka:      kafka,
		repository: repository,
	}
}

// GetAuthBySessionID returns the auth with the given session id.
func (s *AuthService) GetAuthBySessionID(sessionID uint64) (*model.Auth, error) {
	session, err := s.repository.GetSessionByID(sessionID)
	if err != nil {
		return nil, err
	}

	if session == nil || session.DeletedAt.Valid {
		slog.Info("invalid session", "sessionID", sessionID)
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

	hashed, err := s.hashPassword(password)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to hash password: %v", err))
	}

	auth, err := s.repository.CreateUserWithSession(username, hashed, role)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create user with session: %v", err))
	}

	registeredEvent := micro.NewEvent(events.AuthUserRegisteredEventType, UserRegisteredEvent{
		ID:       dto.IDToDTO(auth.User.ID),
		Username: auth.User.Username,
		Role:     dto.RoleToDTO(auth.User.Role),
	})
	err = s.kafka.Producer().Send(events.AuthTopic, registeredEvent.Payload.ID, registeredEvent)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	loggedInEvent := micro.NewEvent(events.AuthUserLoggedInEventType, events.UserLoggedInEvent{
		ID:        dto.IDToDTO(auth.User.ID),
		Username:  auth.User.Username,
		Timestamp: auth.User.LastLogin.Time,
	})
	err = s.kafka.Producer().Send(events.AuthTopic, loggedInEvent.Payload.ID, loggedInEvent)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
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

	if !s.verifyPassword(password, user.Password) {
		return nil, micro.NewBadRequestError("invalid username or password")
	}

	session, err := s.repository.CreateSessionAndUpdateLastLogin(user)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create session: %v", err))
	}

	event := micro.NewEvent(events.AuthUserLoggedInEventType, events.UserLoggedInEvent{
		ID:        dto.IDToDTO(user.ID),
		Username:  user.Username,
		Timestamp: user.LastLogin.Time,
	})
	err = s.kafka.Producer().Send(events.AuthTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return &model.Auth{
		User:    user,
		Session: session,
	}, nil
}

func (s *AuthService) TerminateSession(sessionID uint64) error {
	return s.repository.TerminateSession(sessionID)
}

// hashPassword hashes the given password using bcrypt.
func (s *AuthService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// verifyPassword verifies if the given password matches the stored hash.
func (s *AuthService) verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
