package in

import (
	"github.com/llamadeus/ebike3/packages/auth/domain/port/in"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
)

func MakeLogoutHandler(authService in.AuthService) micro.HTTPHandler {
	type output struct {
		Status string `json:"status"`
	}

	return micro.MakeHandler[any, output](func(ctx micro.Context[any]) (*output, error) {
		sessionID := ctx.Header().Get("X-Session-ID")
		err := authService.TerminateSession(sessionID)
		if err != nil {
			return nil, err
		}

		return &output{
			Status: "success",
		}, nil
	})
}
