package dto

import (
	"api-gw/internal/domain"
)

type UserData struct {
	LoginPasswordCollection []domain.ProtectedItem `json:"login_password_collection"`
	CreditCardCollection    []domain.ProtectedItem `json:"credit_card_collection"`
	TextRecordCollection    []domain.ProtectedItem `json:"text_record_collection"`
	BinaryRecordCollection  []domain.ProtectedItem `json:"binary_record_collection"`
}
