package in

import (
	"github.com/llamadeus/ebike3/packages/auth/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/in"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
)

func MakeAuthHandler(authService in.AuthService) micro.HTTPHandler {
	return micro.MakeHandler(func(ctx micro.Context[any, any]) (*dto.UserDTO, error) {
		sessionID, err := dto.IDFromDTO(ctx.Header().Get("X-Session-ID"))
		if err != nil {
			return nil, micro.NewBadRequestError("invalid session id")
		}

		auth, err := authService.GetAuthBySessionID(sessionID)
		if auth == nil {
			return nil, err
		}

		return dto.UserToDTO(auth.User, auth.Session), nil
	})
}
