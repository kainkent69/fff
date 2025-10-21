package user

import "errors"

var (
	// email
	InvalidEmailErr  = errors.New("Invalid Email.")
	EmailNotFoundErr = errors.New("Email is not Found.")

	// password
	PasswordIsWeak  = errors.New("Weak Password")
	InvalidPassword = errors.New("Invalid Password")
)

func Login(data RegisterAuth) error {

	return nil
}

func Signup(data RegisterAuth) error {

	return nil
}
