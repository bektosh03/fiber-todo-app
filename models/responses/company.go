package responses

type CompanyResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Owner        string `json:"owner"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
