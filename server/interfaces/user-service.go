package interfaces

import "server/dtos"

type IUserService interface {
	CreateUser(user dtos.CreateUser) dtos.CreatedUser
}
