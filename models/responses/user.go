package responses

type User struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Age          uint8  `json:"age"`
	Gender       string `json:"gender"`
	ProfilePhoto string `json:"profile_photo"`
}
