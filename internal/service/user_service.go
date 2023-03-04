package service

import (
	"template/internal/interfaces"
	"template/internal/model"
)

type UserServiceImpl struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) interfaces.UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (s *UserServiceImpl) CreateUser(name, email, password string) (*model.User, error) {
	user := &model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	user, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) GetUser(id int) (*model.User, error) {
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
