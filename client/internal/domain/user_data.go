package domain

type UserData struct {
	LoginPasswordCollection []ProtectedItem `json:"login_password_collection"`
	CreditCardCollection    []ProtectedItem `json:"credit_card_collection"`
	TextRecordCollection    []ProtectedItem `json:"text_record_collection"`
	BinaryRecordCollection  []ProtectedItem `json:"binary_record_collection"`
}
