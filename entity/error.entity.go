package entity

import "errors"

var (
	// ErrNotFound define error if record not found on database
	ErrNotFound = NewError("E-1000", "record not found")
	// ErrUserNotFound define error if auth not found on database
	ErrUserNotFound = NewError("E-1001", "auth not found")
	// ErrOTPMismatch define error if otp mismatch
	ErrOTPMismatch = NewError("E-1002", "OTP mismatch")
	// ErrWrongPasswordConfirmation define error if password confirmation is wrong
	ErrWrongPasswordConfirmation = NewError("E-2001", "Password confirmation is not correct")
	// ErrPasswordMismatch define error if password mismatch
	ErrPasswordMismatch = NewError("E-2001", "Username or password is not correct")
	// ErrOldPasswordMismatch define error if old password mismatch
	ErrOldPasswordMismatch = NewError("E-2002", "Your old password is not correct")
)

// Error define error of gam processor
type Error struct {
	Code    string
	Message string
	Error   error
}

// NewError define new gam processor error
func NewError(
	code,
	message string,
) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Error:   errors.New(message),
	}
}
