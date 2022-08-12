package grpc

import (
	service "storage/internal/application/service"
	pb "storage/pkg/storage_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type StorageHandler struct {
	logger                        Logger
	userService                   service.UserService
	userMetadataFromContextGetter service.UserMetadataFromContextGetter
	loginPasswordService          service.LoginPassword
	creditCardService             service.CreditCard
	textRecordService             service.TextRecord
	binaryRecordService           service.BinaryRecord
	pb.UnimplementedStorageServer
}

func GetStorageHandler(
	logger Logger,
	userService service.UserService,
	userMetadataFromContextGetter service.UserMetadataFromContextGetter,
	loginPasswordService service.LoginPassword,
	creditCardService service.CreditCard,
	textRecordService service.TextRecord,
	binaryRecordService service.BinaryRecord,
) *StorageHandler {
	return &StorageHandler{
		logger:                        logger,
		userService:                   userService,
		userMetadataFromContextGetter: userMetadataFromContextGetter,
		loginPasswordService:          loginPasswordService,
		creditCardService:             creditCardService,
		textRecordService:             textRecordService,
		binaryRecordService:           binaryRecordService,
	}
}
