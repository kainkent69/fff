package user

import (
	"encoding/json"
	"errors"
)

// the user informations.
type User struct {
	// accont information
	Fname string `json:"name.first"`
	Lname string `json:"name.last"`
	// username and email.
	UserName string `json:"username"`
	Email    string `json:"email"`
	// password
	Password string `json:"password"`

	//  coins and shards.
	Shards string `json:"resource.shards"`
	Coins  string `json:"resource.coins"`

	// tax and others
	Level uint `json:"player.level"`
	Tax   uint `json:"player.tax"`
}

// For LogIn

type RegisterAuth struct {
	EmailUsername string `json:"auth.emailOrUsername"`
	Password      string `json:"auth.password"`
}

// the information for registering
type RegisterInfo struct {
	RegisterAuth
	Fname string `json:"name.first"`
	Lname string `json:"name.last"`
}

// convert string to User
func (u User) FromJson(user string) (*User, error) {
	newUser := new(User)
	err := json.Unmarshal([]byte(user), newUser)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// convert a user back json

func (u *User) ToJson() ([]byte, error) {
	data, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return data, nil

}
