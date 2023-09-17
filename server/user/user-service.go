package user

type UserService struct{}

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
		Password:       user.Password,
	}
	return returnValue
}
