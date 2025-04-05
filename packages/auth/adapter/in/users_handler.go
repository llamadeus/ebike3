package in

import (
	"github.com/llamadeus/ebike3/packages/auth/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/auth/domain/port/in"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
)

func MakeUsersHandler(authService in.AuthService) micro.HTTPHandler {
	return micro.MakeHandler(func(ctx micro.Context[any, any]) ([]*dto.UserDTO, error) {
		users, err := authService.GetUsers()
		if users == nil {
			return nil, err
		}

		dtos := make([]*dto.UserDTO, len(users))
		for i, user := range users {
			dtos[i] = dto.UserToDTO(user, nil)
		}

		return dtos, nil
	})
}
