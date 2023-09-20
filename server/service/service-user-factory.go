package service

import "server/interfaces"

type ServiceUserFactory struct {
	userService interfaces.IUserService
}

func (s *ServiceUserFactory) GetUserService(repository interfaces.IUserRepository) interfaces.IUserService {
	if s.userService == nil {
		s.userService = &userService{UserRepository: repository}
	}
	return s.userService
}
