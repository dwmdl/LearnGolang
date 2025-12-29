package di

import "api/internal/user"

type IStatRepository interface {
	AddDirection(linkId uint)
}

type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
}
