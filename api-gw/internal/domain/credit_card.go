package domain

type CreditCard struct {
	Id           string `json:"id"`
	Note         string `json:"note"`
	Name         string `json:"name"`
	Number       string `json:"number"`
	Expiration   string `json:"expiration"`
	Cvv          string `json:"cvv"`
	Owner        string `json:"owner"`
	UserId       string `json:"user_id"`
	CreatedAt    string `json:"created_at"`
	LastAccessAt string `json:"last_access_at"`
}
