package user

import "fmt"

type UserService struct {
	UserRepository IUserRepository
}

type IUserService interface {
	CreateUser(CreateUser) CreateUser
}

func (u *UserService) CreateUser(user *CreateUser) CreatedUser {
	returnValue := CreatedUser{
		Id:             "1",
		Name:           user.Name,
		DocumentNumber: user.DocumentNumber,
		DocumentType:   user.DocumentType,
		Email:          user.Email,
	}
	u.UserRepository.createUser(*user)
	fmt.Println(returnValue.getInternalId())
	return returnValue
}
