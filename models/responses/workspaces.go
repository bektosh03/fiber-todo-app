package responses

type Workspace struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at,omitempty"`
}

type WorkspacesResponse struct {
	Workspaces []Workspace `json:"workspaces"`
	Count      uint64      `json:"count"`
	Page       uint64      `json:"page"`
}
