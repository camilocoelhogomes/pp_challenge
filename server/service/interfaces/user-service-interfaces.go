package service

import dtos "server/dtos/user"

type IUserService interface {
	CreateUser(user dtos.CreateUser) dtos.CreatedUser
}
