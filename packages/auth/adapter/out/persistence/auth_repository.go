package persistence

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/llamadeus/ebike3/packages/auth/domain/model"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/out"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/database"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/utils"
	"time"
)

type AuthRepository struct {
	db        *sqlx.DB
	snowflake *utils.SnowflakeGenerator
}

var _ out.AuthRepository = (*AuthRepository)(nil)

func NewAuthRepository(db *sqlx.DB, snowflake *utils.SnowflakeGenerator) *AuthRepository {
	return &AuthRepository{db: db, snowflake: snowflake}
}

func (r *AuthRepository) CreateUserWithSession(username string, password string, role model.UserRole) (*model.Auth, error) {
	userID, err := r.snowflake.Generate()
	if err != nil {
		return nil, err
	}

	sessionID, err := r.snowflake.Generate()
	if err != nil {
		return nil, err
	}

	err = database.RunInTx(r.db, func(tx *sqlx.Tx) error {
		_, err = tx.NamedExec("INSERT INTO users (id, username, password, role, last_login) VALUES (:id, :username, :password, :role, :last_login)", map[string]any{
			"id":         userID,
			"username":   username,
			"password":   password,
			"role":       role,
			"last_login": time.Now(),
		})
		if err != nil {
			return err
		}

		_, err = tx.NamedExec("INSERT INTO sessions (id, user_id) VALUES (:id, :user_id)", map[string]any{
			"id":      sessionID,
			"user_id": userID,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	user, err := r.GetUserByID(userID)
	if user == nil {
		return nil, err
	}

	session, err := r.GetSessionByID(sessionID)
	if session == nil {
		return nil, err
	}

	return &model.Auth{
		User:    user,
		Session: session,
	}, nil
}

func (r *AuthRepository) CreateSessionAndUpdateLastLogin(user *model.User) (*model.Session, error) {
	sessionID, err := r.snowflake.Generate()
	if err != nil {
		return nil, err
	}

	err = database.RunInTx(r.db, func(tx *sqlx.Tx) error {
		_, err = r.db.NamedExec("INSERT INTO sessions (id, user_id) VALUES (:id, :user_id)", map[string]any{
			"id":      sessionID,
			"user_id": user.ID,
		})
		if err != nil {
			return err
		}

		_, err = tx.NamedExec("UPDATE users SET last_login=:last_login WHERE id=:id", map[string]any{
			"id":         user.ID,
			"last_login": time.Now(),
		})
		if err != nil {
			return err
		}

		return nil
	})

	return r.GetSessionByID(sessionID)
}

func (r *AuthRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := r.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *AuthRepository) GetUserByID(id uint64) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id=$1 LIMIT 1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE username=$1 LIMIT 1", username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) GetSessionByID(id uint64) (*model.Session, error) {
	var session model.Session
	err := r.db.Get(&session, "SELECT * FROM sessions WHERE id=$1 LIMIT 1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &session, nil
}

func (r *AuthRepository) TerminateSession(id uint64) error {
	_, err := r.db.Exec("UPDATE sessions SET deleted_at=NOW() WHERE id=$1", id)

	return err
}
