package services

import "errors"

var (
	ErrFailedToParseRequest = errors.New("failed to parse request")
	ErrDuplicateName        = errors.New("name already exists")
	ErrDuplicateEmail       = errors.New("email already exists")
	ErrPasswordMismatch     = errors.New("password and confirm password do not match")
	ErrFailedToRegisterUser = errors.New("failed to register user")
	ErrFailedToHashPassword = errors.New("failed to hash password")
	ErrUnexpectedError      = errors.New("unexpected error")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrInvalidParameters    = errors.New("invalid parameters")
	ErrInvalidTaskID        = errors.New("invalid task id")
	ErrRecordNotFound       = errors.New("record not found")
)
