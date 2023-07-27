package services

import (
	userModel "hexagonal-architecture/app/users/domain/models"
	userInterface "hexagonal-architecture/app/users/domain/repositories"
)

type Service struct {
	userRepository userInterface.IUser
}

func NewService(userRepository userInterface.IUser) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) Create(user *userModel.User) string {
	return s.userRepository.Create(user)
}

func (s *Service) Get() []*userModel.User {
	return s.userRepository.Get()
}
