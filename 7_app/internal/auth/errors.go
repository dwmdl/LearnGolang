package auth

import "errors"

var (
	ErrUserExists       = errors.New("user exists")
	ErrUserDoesntExists = errors.New("user doesn't exists")
	ErrWrongCredentials = errors.New("incorrect credentials")
)
