package interfaces

import "server/dtos"

type IUserRepository interface {
	CreateUser(user dtos.CreateUser) (externalId *string, err *error)
}
