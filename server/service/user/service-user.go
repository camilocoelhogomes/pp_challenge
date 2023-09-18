package service

import (
	"fmt"
	dtos "server/dtos/user"
	"server/interfaces"
)

type UserService struct {
	UserRepository interfaces.IUserRepository
}

func (u *UserService) CreateUser(user *dtos.CreateUser) dtos.CreatedUser {
	fmt.Println("userService memory address => %s", &u)
	returnValue := dtos.CreatedUser{
		Name:           user.Name,
		DocumentNumber: user.DocumentNumber,
		DocumentType:   user.DocumentType,
		Email:          user.Email,
	}
	result, _ := u.UserRepository.CreateUser(*user)
	returnValue.Id = *result

	return returnValue
}
