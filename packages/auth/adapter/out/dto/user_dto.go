package dto

import (
	"github.com/llamadeus/ebike3/packages/auth/domain/model"
	"time"
)

type UserDTO struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Role      string `json:"role,omitempty"`
	SessionID string `json:"sessionId,omitempty"`
	LastLogin string `json:"lastLogin,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

func UserToDTO(user *model.User, session *model.Session) *UserDTO {
	return &UserDTO{
		ID:        IDToDTO(user.ID),
		Username:  user.Username,
		Role:      RoleToDTO(user.Role),
		LastLogin: user.LastLogin.Time.Format(time.RFC3339),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		SessionID: IDToDTO(session.ID),
	}
}

func RoleToDTO(role model.UserRole) string {
	return string(role)
}
