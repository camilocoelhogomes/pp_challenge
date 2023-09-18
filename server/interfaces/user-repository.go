package interfaces

import dtos "server/dtos/user"

type IUserRepository interface {
	CreateUser(user dtos.CreateUser) (externalId *string, err *error)
}
