package user

import "errors"

// Validate register

func Validate(u *RegisterInfo) (bool, []byte, error) {
	found := new(User)
	if ok := len(found.Lname) >= 2 || len(found.Fname) >= 3; !ok {
		return false, []byte(`user.Validate: Invalid Name`), errors.New("error")
	}

	return true, []byte("user.Validate: Success"), nil

}
