package service

import (
	"context"
	"google.golang.org/grpc/metadata"
	"user-data-manager/internal/infrastructure/clients/grpc"
	storagePb "user-data-manager/pkg/storage_grpc"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

type UserData struct {
	userIdFromContextGetter UserIdFromContextGetter
	storageClient           grpc.StorageClient
}

func GetUserDataService(userIdFromContextGetter UserIdFromContextGetter, storageClient grpc.StorageClient) UserData {
	return UserData{
		userIdFromContextGetter: userIdFromContextGetter,
		storageClient:           storageClient,
	}
}

func (s UserData) GetLoginPasswordById(id string, ctx context.Context) (pb.LoginPassword, error) {
	var loginPassword pb.LoginPassword

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).GetLoginPasswordById(
		ctx,
		&storagePb.GetLoginPasswordByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		return loginPassword, err
	}

	loginPassword.Id = storageResponse.LoginPassword.Id
	loginPassword.Login = storageResponse.LoginPassword.Login
	loginPassword.Password = storageResponse.LoginPassword.Password
	loginPassword.Name = storageResponse.LoginPassword.Name
	loginPassword.Note = storageResponse.LoginPassword.Note
	loginPassword.LastAccess = storageResponse.LoginPassword.LastAccess
	loginPassword.CreatedDate = storageResponse.LoginPassword.CreatedDate

	return loginPassword, nil
}
