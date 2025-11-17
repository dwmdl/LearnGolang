package auth

import (
	"api/internal/user"

	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	userModel := &user.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	}

	_, err = service.UserRepository.Create(userModel)
	if err != nil {
		return "", err
	}
	return userModel.Email, nil
}
