package in

import (
	"github.com/llamadeus/ebike3/packages/auth/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/in"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
)

func MakeLogoutHandler(authService in.AuthService) micro.HTTPHandler {
	type output struct {
		Status string `json:"status"`
	}

	return micro.MakeHandler(func(ctx micro.Context[any, any]) (*output, error) {
		sessionID, err := dto.IDFromDTO(ctx.Header().Get("X-Session-ID"))
		if err != nil {
			return nil, micro.NewBadRequestError("invalid session id")
		}

		err = authService.TerminateSession(sessionID)
		if err != nil {
			return nil, err
		}

		return &output{
			Status: "success",
		}, nil
	})
}
