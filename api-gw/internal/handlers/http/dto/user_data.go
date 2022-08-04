package dto

import (
	"api-gw/pkg/user_data_manager_grpc"
)

type UserData struct {
	LoginPasswordCollection []*user_data_manager_grpc.ProtectedItem `json:"login_password_collection"`
	CreditCardCollection    []*user_data_manager_grpc.ProtectedItem `json:"credit_card_collection"`
	TextRecordCollection    []*user_data_manager_grpc.ProtectedItem `json:"text_record_collection"`
	BinaryRecordCollection  []*user_data_manager_grpc.ProtectedItem `json:"binary_record_collection"`
}
