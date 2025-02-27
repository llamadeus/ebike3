package in

import (
	"github.com/llamadeus/ebike3/packages/auth/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/auth/domain/model"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/in"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
)

func MakeRegisterHandler(authService in.AuthService) micro.HTTPHandler {
	type input struct {
		Username string         `json:"username" validate:"required"`
		Password string         `json:"password" validate:"required"`
		Role     model.UserRole `json:"role" validate:"required,oneof=ADMIN CUSTOMER"`
	}

	return micro.MakeHandler(func(ctx micro.Context[any, input]) (*dto.UserDTO, error) {
		auth, err := authService.RegisterUser(ctx.Input().Username, ctx.Input().Password, ctx.Input().Role)
		if auth == nil {
			return nil, err
		}

		return dto.UserToDTO(auth.User, auth.Session), nil
	})
}
