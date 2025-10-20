package auth

import (
	"api/internal/user"
)

type Service struct {
	UserRepository *user.Repository
}

func NewService(userRepository *user.Repository) *Service {
	return &Service{UserRepository: userRepository}
}

func (service *Service) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.GetByEmail(email)
	if existedUser != nil {
		return "", ErrUserExists
	}

	userModel := &user.User{
		Email:    email,
		Name:     name,
		Password: "",
	}

	_, err := service.UserRepository.Create(userModel)
	if err != nil {
		return "", err
	}
	return userModel.Email, nil
}
