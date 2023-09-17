package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	Db gorm.DB
}

type IUserRepository interface {
	createUser(user CreateUser) (externalId *string, err *error)
}

func (u *UserRepository) createUser(user CreateUser) (externalId *string, err *error) {
	createdUser := u.Db.Create(&CreatedUser{
		Name:           user.Name,
		DocumentNumber: user.DocumentNumber,
		DocumentType:   user.DocumentType,
		Email:          user.Email,
	})
	if createdUser.Error != nil {
		return nil, &createdUser.Error
	}
	returnValue := "created"
	return &returnValue, nil
}
