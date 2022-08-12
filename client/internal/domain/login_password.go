package domain

type LoginPassword struct {
	Id           string `json:"id"`
	Note         string `json:"note"`
	Name         string `json:"name"`
	Login        string `json:"login"`
	Password     string `json:"password"`
	UserId       string `json:"user_id"`
	CreatedAt    string `json:"created_at"`
	LastAccessAt string `json:"last_access_at"`
}
