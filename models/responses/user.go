package responses

type User struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Age          uint8  `json:"age"`
	Gender       string `json:"gender"`
	ProfilePhoto string `json:"profile_photo"`
	Username     string `json:"username"`
	CompanyID    string `json:"company_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GetCompanyUsers struct {
	CompanyID   string `json:"company_id"`
	CompanyName string `json:"company_name"`
	Users       []User `json:"users"`
	Page        int    `json:"page"`
	Count       int    `json:"count"`
}
