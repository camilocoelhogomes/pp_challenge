package user

import "fmt"

type UserService struct {
	UserRepository IUserRepository
}

type IUserService interface {
	CreateUser(user CreateUser) CreateUser
}

func (u *UserService) CreateUser(user *CreateUser) CreatedUser {
	fmt.Println("userService memory address => %s", &u)
	returnValue := CreatedUser{
		Name:           user.Name,
		DocumentNumber: user.DocumentNumber,
		DocumentType:   user.DocumentType,
		Email:          user.Email,
	}
	result, _ := u.UserRepository.createUser(*user)

	returnValue.Id = *result
	fmt.Println(returnValue.getInternalId())
	return returnValue
}
