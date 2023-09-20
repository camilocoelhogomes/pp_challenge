package service

import (
	"fmt"
	"server/dtos"
	"server/interfaces"
)

type userService struct {
	UserRepository interfaces.IUserRepository
}

func (u *userService) CreateUser(user dtos.CreateUser) dtos.CreatedUser {
	fmt.Println("userService memory address => %s", &u)
	returnValue := dtos.CreatedUser{
		Name:           user.Name,
		DocumentNumber: user.DocumentNumber,
		DocumentType:   user.DocumentType,
		Email:          user.Email,
	}
	result, _ := u.UserRepository.CreateUser(user)
	returnValue.Id = *result

	return returnValue
}
