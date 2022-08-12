package domain

type ProtectedItem struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	CreatedAt    string `json:"created_at"`
	LastAccessAt string `json:"last_access_at"`
}
