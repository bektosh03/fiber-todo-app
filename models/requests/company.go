package requests

type CreateCompany struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}