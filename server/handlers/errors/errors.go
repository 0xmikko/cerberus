package errors

import "errors"

var (
	ErrorWrongRequest                    = errors.New("Wrong request")
	ErrorUserNotFoundOrIncorrectPassword = errors.New("User not found or incorrect password")
	ErrorInternalServerError             = errors.New("Internal Server error")
)
