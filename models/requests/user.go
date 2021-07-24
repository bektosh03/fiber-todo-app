package requests

import (
	"github.com/bektosh/fiber-app/api/validators"
)

type CreateUser struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Age          uint8  `json:"age"`
	Gender       string `json:"gender"`
	ProfilePhoto string `json:"profile_photo"`
	CompanyID    string `json:"company_id"`
}

func (cu CreateUser) Validate() (ok bool) {
	ok = validators.ValidateEmail(cu.Email)
	if !ok {
		return false
	}
	ok = validators.ValidatePassword(cu.Password)
	if !ok {
		return false
	}
	return ok
}
