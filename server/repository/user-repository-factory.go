package repository

import (
	"server/interfaces"

	"gorm.io/gorm"
)

type UserRepositoryFactory struct {
	userRepository interfaces.IUserRepository
}

func (u *UserRepositoryFactory) CreateUserRespotirory(db gorm.DB) interfaces.IUserRepository {
	if u.userRepository == nil {
		u.userRepository = &userRepository{db: db}
	}
	return u.userRepository
}
