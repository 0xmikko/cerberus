package errors

import "errors"

var (
	ErrorUserNotFound           = errors.New("User Not Found")
	ErrorUserWithUsernameExists = errors.New("User with this username exists")
	ErrorRefreshTokenIsNotValid = errors.New("Refresh token is not valid")
	ErrorSignUp                 = errors.New("SignUpError")
	ErrorCantCreateAccounts     = errors.New("Cant create account")
)

