package in

import (
	"github.com/llamadeus/ebike3/packages/auth/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/in"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
)

func MakeLoginHandler(authService in.AuthService) micro.HTTPHandler {
	type input struct {
		Username string `json:"username" validate:"required,min=3"`
		Password string `json:"password" validate:"required,min=8"`
	}

	return micro.MakeHandler(func(ctx micro.Context[any, input]) (*dto.UserDTO, error) {
		auth, err := authService.LoginUser(ctx.Input().Username, ctx.Input().Password)
		if auth == nil {
			return nil, err
		}

		return dto.UserToDTO(auth.User, auth.Session), nil
	})
}
