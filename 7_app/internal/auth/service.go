package auth

import (
	"api/internal/user"
	"api/pkg/di"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepository di.IUserRepository
}

func NewService(userRepository di.IUserRepository) *Service {
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

	newUser, err := service.UserRepository.Create(userModel)
	if err != nil {
		return "", err
	}

	return newUser.Email, nil
}

func (service *Service) Login(email, password string) (string, error) {
	foundUser, _ := service.UserRepository.GetByEmail(email)
	if foundUser == nil {
		return "", ErrUserDoesntExists
	}

	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return "", ErrWrongCredentials
	}

	return foundUser.Email, nil
}
