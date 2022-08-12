package domain

type TextRecord struct {
	Id           string `json:"id"`
	Note         string `json:"note"`
	Name         string `json:"name"`
	Text         string `json:"text"`
	UserId       string `json:"user_id"`
	CreatedAt    string `json:"created_at"`
	LastAccessAt string `json:"last_access_at"`
}
