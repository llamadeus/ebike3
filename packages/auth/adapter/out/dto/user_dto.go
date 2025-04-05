package dto

import (
	"github.com/guregu/null/v5"
	"github.com/llamadeus/ebike3/packages/auth/domain/model"
	"time"
)

type UserDTO struct {
	ID        string      `json:"id"`
	Username  string      `json:"username"`
	Role      string      `json:"role"`
	SessionID null.String `json:"sessionId"`
	LastLogin string      `json:"lastLogin"`
	CreatedAt string      `json:"createdAt"`
	UpdatedAt string      `json:"updatedAt"`
}

func UserToDTO(user *model.User, session *model.Session) *UserDTO {
	var sessionID null.String
	if session != nil {
		sessionID.SetValid(IDToDTO(session.ID))
	}

	return &UserDTO{
		ID:        IDToDTO(user.ID),
		Username:  user.Username,
		Role:      RoleToDTO(user.Role),
		LastLogin: user.LastLogin.Time.Format(time.RFC3339),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		SessionID: sessionID,
	}
}

func RoleToDTO(role model.UserRole) string {
	return string(role)
}
