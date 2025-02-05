package dto

import (
	"github.com/llamadeus/ebike3/packages/auth/domain/model"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/utils"
	"time"
)

type UserDTO struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	SessionID string `json:"sessionId"`
	LastLogin string `json:"lastLogin"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func UserToDTO(user *model.User, session *model.Session) *UserDTO {
	return &UserDTO{
		ID:        string(utils.Base62.FormatUint(user.ID)),
		Username:  user.Username,
		Role:      string(user.Role),
		LastLogin: user.LastLogin.Time.Format(time.RFC3339),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		SessionID: session.SessionID,
	}
}
