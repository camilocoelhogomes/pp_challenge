package user

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

type IUserRepository interface {
	createUser(user CreateUser) (externalId *string, err *error)
}

func (u *UserRepository) createUser(user CreateUser) (externalId *string, err *error) {
	fmt.Println("UserRepository memory address => %s", &u)
	fmt.Println("Db memory address => %s", &u.Db)
	userToBeCreated := CreatedUser{
		Name:           user.Name,
		DocumentNumber: user.DocumentNumber,
		DocumentType:   user.DocumentType,
		Email:          user.Email,
	}
	createdUser := u.Db.Create(&userToBeCreated)
	if createdUser.Error != nil {
		return nil, &createdUser.Error
	}
	return &userToBeCreated.Id, nil
}
