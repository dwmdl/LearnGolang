package auth_test

import (
	"api/internal/auth"
	"api/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.ru",
	}, nil
}

func (repo *MockUserRepository) GetByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestServiceRegisterSuccess(t *testing.T) {
	const originEmail = "a@a.ru"

	authService := auth.NewService(&MockUserRepository{})
	email, err := authService.Register(originEmail, "123", "TestAuthService")
	if err != nil {
		t.Fatal(err.Error())
	}
	if email != originEmail {
		t.Fatalf("Email %s does not match %s", email, originEmail)
	}
}
