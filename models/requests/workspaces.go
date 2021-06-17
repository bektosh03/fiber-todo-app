package requests

type CreateWorkspace struct {
	Name string `json:"name"`
}

type UpdateWorkspace struct {
	ID string `json:"id"`
	Name string `json:"name"`
}
