package domain

type LoginPassword struct {
	Id           string
	Note         string
	Name         string
	Login        string
	Password     string
	UserId       string
	CreatedAt    string `json:"created_at"`
	LastAccessAt string `json:"last_access_at"`
}
