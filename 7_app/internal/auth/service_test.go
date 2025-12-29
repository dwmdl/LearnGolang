package auth

import (
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

func TestRegisterSuccess(t *testing.T) {
	const originEmail = "a@a.ru"

	authService := NewService(&MockUserRepository{})
	email, err := authService.Register(originEmail, "123", "TestAuthService")
	if err != nil {
		t.Fatal(err.Error())
	}
	if email != originEmail {
		t.Fatalf("Email %s does not match %s", email, originEmail)
	}
}
