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

func (s UserData) AddLoginPassword(request pb.CreateLoginPassword, ctx context.Context) (pb.LoginPassword, error) {
	var loginPassword pb.LoginPassword

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).CreateLoginPassword(
		ctx,
		&storagePb.CreateLoginPasswordRequest{
			CreateLoginPassword: &storagePb.CreateLoginPassword{
				Name:     request.Name,
				Login:    request.Login,
				Password: request.Password,
				Note:     request.Note,
				UserId:   s.userIdFromContextGetter.getUserIdFromContext(ctx),
			},
		},
	)

	if err != nil {
		return loginPassword, err
	}

	loginPassword.Id = storageResponse.LoginPassword.Id
	loginPassword.Name = storageResponse.LoginPassword.Name
	loginPassword.Login = storageResponse.LoginPassword.Login
	loginPassword.Password = storageResponse.LoginPassword.Password
	loginPassword.Note = storageResponse.LoginPassword.Note
	loginPassword.LastAccess = storageResponse.LoginPassword.LastAccess
	loginPassword.CreatedDate = storageResponse.LoginPassword.CreatedDate

	return loginPassword, nil
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

func (s UserData) GetLoginPasswordList(ctx context.Context) ([]pb.LoginPassword, error) {
	var loginPasswordList []pb.LoginPassword

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).GetLoginPasswordList(
		ctx,
		&storagePb.GetLoginPasswordListRequest{},
	)

	if err != nil {
		return nil, err
	}

	var loginPassword pb.LoginPassword

	for _, loginPasswordResponse := range storageResponse.LoginPasswordList {
		loginPassword.Id = loginPasswordResponse.Id
		loginPassword.Login = loginPasswordResponse.Login
		loginPassword.Password = loginPasswordResponse.Password
		loginPassword.Name = loginPasswordResponse.Name
		loginPassword.Note = loginPasswordResponse.Note
		loginPassword.LastAccess = loginPasswordResponse.LastAccess
		loginPassword.CreatedDate = loginPasswordResponse.CreatedDate

		loginPasswordList = append(loginPasswordList, loginPassword)
	}

	return loginPasswordList, nil
}

func (s UserData) DeleteLoginPasswordById(id string, ctx context.Context) error {
	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	_, err := (*s.storageClient.GetClient()).DeleteLoginPasswordById(
		ctx,
		&storagePb.DeleteLoginPasswordByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
