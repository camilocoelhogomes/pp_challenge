package repository

import (
	"server/dtos"

	"gorm.io/gorm"
)

type userRepository struct {
	db gorm.DB
}

func (u *userRepository) SetDb(db gorm.DB) {
	u.db = db
}

func (u *userRepository) CreateUser(user dtos.CreateUser) (externalId *string, err *error) {
	userToBeCreated := dtos.CreatedUser{
		Name:           user.Name,
		DocumentNumber: user.DocumentNumber,
		DocumentType:   user.DocumentType,
		Email:          user.Email,
	}
	createdUser := u.db.Create(&userToBeCreated)
	if createdUser.Error != nil {
		return nil, &createdUser.Error
	}
	return &userToBeCreated.Id, nil
}
