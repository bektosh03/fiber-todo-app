package requests

import (
	"github.com/bektosh/fiber-app/api/validators"
)

type CreateUser struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Age          uint8  `json:"age"`
	Gender       string `json:"gender"`
	ProfilePhoto string `json:"profile_photo"`
}

func (cu CreateUser) Validate() (ok bool, err error) {
	ok, err = validators.ValidateEmail(cu.Email)
	if err != nil || !ok {
		return false, err
	}
	ok, err = validators.ValidatePassword(cu.Password)
	if err != nil || !ok {
		return false, err
	}
	return ok, nil
}
