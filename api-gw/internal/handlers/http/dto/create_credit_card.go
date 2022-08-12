package dto

type CreateCreditCard struct {
	Name       string
	Number     string
	Expiration string
	Cvv        string
	Owner      string
	Note       string
}
