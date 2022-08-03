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

func (s UserData) UpdateLoginPassword(request pb.LoginPassword, ctx context.Context) (pb.LoginPassword, error) {
	var loginPassword pb.LoginPassword

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).UpdateLoginPasswordById(
		ctx,
		&storagePb.UpdateLoginPasswordByIdRequest{
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

/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////

func (s UserData) AddCreditCard(request pb.CreateCreditCard, ctx context.Context) (pb.CreditCard, error) {
	var creditCard pb.CreditCard

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).CreateCreditCard(
		ctx,
		&storagePb.CreateCreditCardRequest{
			CreateCreditCard: &storagePb.CreateCreditCard{
				Name:       request.Name,
				Number:     request.Number,
				Expiration: request.Expiration,
				Cvv:        request.Cvv,
				Owner:      request.Owner,
				Note:       request.Note,
				UserId:     s.userIdFromContextGetter.getUserIdFromContext(ctx),
			},
		},
	)

	if err != nil {
		return creditCard, err
	}

	creditCard.Id = storageResponse.CreditCard.Id
	creditCard.Name = storageResponse.CreditCard.Name
	creditCard.Number = storageResponse.CreditCard.Number
	creditCard.Expiration = storageResponse.CreditCard.Expiration
	creditCard.Cvv = storageResponse.CreditCard.Cvv
	creditCard.Owner = storageResponse.CreditCard.Owner
	creditCard.Note = storageResponse.CreditCard.Note
	creditCard.LastAccess = storageResponse.CreditCard.LastAccess
	creditCard.CreatedDate = storageResponse.CreditCard.CreatedDate

	return creditCard, nil
}

func (s UserData) GetCreditCardById(id string, ctx context.Context) (pb.CreditCard, error) {
	var creditCard pb.CreditCard

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).GetCreditCardById(
		ctx,
		&storagePb.GetCreditCardByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		return creditCard, err
	}

	creditCard.Id = storageResponse.CreditCard.Id
	creditCard.Number = storageResponse.CreditCard.Number
	creditCard.Expiration = storageResponse.CreditCard.Expiration
	creditCard.Cvv = storageResponse.CreditCard.Cvv
	creditCard.Owner = storageResponse.CreditCard.Owner
	creditCard.Name = storageResponse.CreditCard.Name
	creditCard.Note = storageResponse.CreditCard.Note
	creditCard.LastAccess = storageResponse.CreditCard.LastAccess
	creditCard.CreatedDate = storageResponse.CreditCard.CreatedDate

	return creditCard, nil
}

func (s UserData) GetCreditCardList(ctx context.Context) ([]pb.CreditCard, error) {
	var creditCardList []pb.CreditCard

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).GetCreditCardList(
		ctx,
		&storagePb.GetCreditCardListRequest{},
	)

	if err != nil {
		return nil, err
	}

	var creditCard pb.CreditCard

	for _, creditCardResponse := range storageResponse.CreditCardList {
		creditCard.Id = creditCardResponse.Id
		creditCard.Number = creditCardResponse.Number
		creditCard.Expiration = creditCardResponse.Expiration
		creditCard.Cvv = creditCardResponse.Cvv
		creditCard.Owner = creditCardResponse.Owner
		creditCard.Name = creditCardResponse.Name
		creditCard.Note = creditCardResponse.Note
		creditCard.LastAccess = creditCardResponse.LastAccess
		creditCard.CreatedDate = creditCardResponse.CreatedDate

		creditCardList = append(creditCardList, creditCard)
	}

	return creditCardList, nil
}

func (s UserData) DeleteCreditCardById(id string, ctx context.Context) error {
	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	_, err := (*s.storageClient.GetClient()).DeleteCreditCardById(
		ctx,
		&storagePb.DeleteCreditCardByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (s UserData) UpdateCreditCard(request pb.CreditCard, ctx context.Context) (pb.CreditCard, error) {
	var creditCard pb.CreditCard

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).UpdateCreditCardById(
		ctx,
		&storagePb.UpdateCreditCardByIdRequest{
			CreateCreditCard: &storagePb.CreateCreditCard{
				Name:       request.Name,
				Number:     request.Number,
				Expiration: request.Expiration,
				Cvv:        request.Cvv,
				Owner:      request.Owner,
				Note:       request.Note,
				UserId:     s.userIdFromContextGetter.getUserIdFromContext(ctx),
			},
		},
	)

	if err != nil {
		return creditCard, err
	}

	creditCard.Id = storageResponse.CreditCard.Id
	creditCard.Name = storageResponse.CreditCard.Name
	creditCard.Number = storageResponse.CreditCard.Number
	creditCard.Expiration = storageResponse.CreditCard.Expiration
	creditCard.Cvv = storageResponse.CreditCard.Cvv
	creditCard.Owner = storageResponse.CreditCard.Owner
	creditCard.Note = storageResponse.CreditCard.Note
	creditCard.LastAccess = storageResponse.CreditCard.LastAccess
	creditCard.CreatedDate = storageResponse.CreditCard.CreatedDate

	return creditCard, nil
}
