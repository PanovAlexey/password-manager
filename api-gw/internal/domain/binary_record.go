package domain

type BinaryRecord struct {
	Id           string `json:"id"`
	Note         string `json:"note"`
	Name         string `json:"name"`
	Binary       string `json:"binary"`
	UserId       string `json:"user_id"`
	CreatedAt    string `json:"created_at"`
	LastAccessAt string `json:"last_access_at"`
}
