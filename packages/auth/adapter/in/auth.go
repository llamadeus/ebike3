package in

import (
	"github.com/llamadeus/ebike3/packages/auth/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/in"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
)

func NewAuthHandler(authService in.AuthService) micro.HTTPHandler {
	return micro.MakeHandler[any, dto.UserDTO](func(ctx micro.Context[any]) (*dto.UserDTO, error) {
		sessionID := ctx.Header().Get("X-Session-ID")
		auth, err := authService.GetAuthBySessionID(sessionID)
		if auth == nil {
			return nil, err
		}

		return dto.UserToDTO(auth.User, auth.Session), nil
	})
}
