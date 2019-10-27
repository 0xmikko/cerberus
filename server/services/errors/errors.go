package errors

import "errors"

var (
	ErrorUserNotFound           = errors.New("User Not Found")
	ErrorUserWithUsernameExists = errors.New("User with this username exists")
	ErrorRefreshTokenIsNotValid = errors.New("Refresh token is not valid")
	ErrorSignUp                 = errors.New("SignUpError")
	ErrorCantCreateAccounts     = errors.New("Cant create account")
	ErrorTransactionNotFound    = errors.New("Transaction not found")
	ErrorCantConfirmCancelled   = errors.New("Cant confirm cancelled transaction")
	ErrorDBError                = errors.New("Database error")
	ErrorHaveNoRights           = errors.New("You have no rights for this operation")
)
