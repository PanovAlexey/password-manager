package service

import (
	"context"
	"google.golang.org/grpc/metadata"
	"user-data-manager/internal/infrastructure/clients/grpc"
	storagePb "user-data-manager/pkg/storage_grpc"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

type UserData struct {
	storageClient grpc.StorageClient
}

func GetUserDataService(storageClient grpc.StorageClient) UserData {
	return UserData{
		storageClient: storageClient,
	}
}

func (s UserData) AddLoginPassword(request pb.CreateLoginPassword, userId string, ctx context.Context) (pb.LoginPassword, error) {
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
				UserId:   userId,
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

func (s UserData) UpdateLoginPassword(request pb.LoginPassword, userId string, ctx context.Context) (pb.LoginPassword, error) {
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
				UserId:   userId,
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

func (s UserData) AddCreditCard(request pb.CreateCreditCard, userId string, ctx context.Context) (pb.CreditCard, error) {
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
				UserId:     userId,
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

func (s UserData) UpdateCreditCard(request pb.CreditCard, userId string, ctx context.Context) (pb.CreditCard, error) {
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
				UserId:     userId,
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

func (s UserData) AddTextRecord(request pb.TextRecord, userId string, ctx context.Context) (pb.TextRecord, error) {
	var textRecord pb.TextRecord

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).CreateTextRecord(
		ctx,
		&storagePb.CreateTextRecordRequest{
			CreateTextRecord: &storagePb.CreateTextRecord{
				Name:   request.Name,
				Text:   request.Text,
				Note:   request.Note,
				UserId: userId,
			},
		},
	)

	if err != nil {
		return textRecord, err
	}

	textRecord.Id = storageResponse.TextRecord.Id
	textRecord.Name = storageResponse.TextRecord.Name
	textRecord.Text = storageResponse.TextRecord.Text
	textRecord.Note = storageResponse.TextRecord.Note
	textRecord.LastAccess = storageResponse.TextRecord.LastAccess
	textRecord.CreatedDate = storageResponse.TextRecord.CreatedDate

	return textRecord, nil
}

func (s UserData) GetTextRecordById(id string, ctx context.Context) (pb.TextRecord, error) {
	var textRecord pb.TextRecord

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).GetTextRecordById(
		ctx,
		&storagePb.GetTextRecordByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		return textRecord, err
	}

	textRecord.Id = storageResponse.TextRecord.Id
	textRecord.Text = storageResponse.TextRecord.Text
	textRecord.Name = storageResponse.TextRecord.Name
	textRecord.Note = storageResponse.TextRecord.Note
	textRecord.LastAccess = storageResponse.TextRecord.LastAccess
	textRecord.CreatedDate = storageResponse.TextRecord.CreatedDate

	return textRecord, nil
}

func (s UserData) GetTextRecordList(ctx context.Context) ([]pb.TextRecord, error) {
	var textRecordList []pb.TextRecord

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).GetTextRecordList(
		ctx,
		&storagePb.GetTextRecordListRequest{},
	)

	if err != nil {
		return nil, err
	}

	var textRecord pb.TextRecord

	for _, textRecordResponse := range storageResponse.TextRecordList {
		textRecord.Id = textRecordResponse.Id
		textRecord.Text = textRecordResponse.Text
		textRecord.Name = textRecordResponse.Name
		textRecord.Note = textRecordResponse.Note
		textRecord.LastAccess = textRecordResponse.LastAccess
		textRecord.CreatedDate = textRecordResponse.CreatedDate

		textRecordList = append(textRecordList, textRecord)
	}

	return textRecordList, nil
}

func (s UserData) DeleteTextRecordById(id string, ctx context.Context) error {
	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	_, err := (*s.storageClient.GetClient()).DeleteTextRecordById(
		ctx,
		&storagePb.DeleteTextRecordByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (s UserData) UpdateTextRecord(request pb.TextRecord, userId string, ctx context.Context) (pb.TextRecord, error) {
	var textRecord pb.TextRecord

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).UpdateTextRecordById(
		ctx,
		&storagePb.UpdateTextRecordByIdRequest{
			CreateTextRecord: &storagePb.CreateTextRecord{
				Name:   request.Name,
				Text:   request.Text,
				Note:   request.Note,
				UserId: userId,
			},
		},
	)

	if err != nil {
		return textRecord, err
	}

	textRecord.Id = storageResponse.TextRecord.Id
	textRecord.Name = storageResponse.TextRecord.Name
	textRecord.Text = storageResponse.TextRecord.Text
	textRecord.Note = storageResponse.TextRecord.Note
	textRecord.LastAccess = storageResponse.TextRecord.LastAccess
	textRecord.CreatedDate = storageResponse.TextRecord.CreatedDate

	return textRecord, nil
}

func (s UserData) AddBinaryRecord(request pb.BinaryRecord, userId string, ctx context.Context) (pb.BinaryRecord, error) {
	var binaryRecord pb.BinaryRecord

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).CreateBinaryRecord(
		ctx,
		&storagePb.CreateBinaryRecordRequest{
			CreateBinaryRecord: &storagePb.CreateBinaryRecord{
				Name:   request.Name,
				Binary: request.Binary,
				Note:   request.Note,
				UserId: userId,
			},
		},
	)

	if err != nil {
		return binaryRecord, err
	}

	binaryRecord.Id = storageResponse.BinaryRecord.Id
	binaryRecord.Name = storageResponse.BinaryRecord.Name
	binaryRecord.Binary = storageResponse.BinaryRecord.Binary
	binaryRecord.Note = storageResponse.BinaryRecord.Note
	binaryRecord.LastAccess = storageResponse.BinaryRecord.LastAccess
	binaryRecord.CreatedDate = storageResponse.BinaryRecord.CreatedDate

	return binaryRecord, nil
}

func (s UserData) GetBinaryRecordById(id string, ctx context.Context) (pb.BinaryRecord, error) {
	var binaryRecord pb.BinaryRecord

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).GetBinaryRecordById(
		ctx,
		&storagePb.GetBinaryRecordByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		return binaryRecord, err
	}

	binaryRecord.Id = storageResponse.BinaryRecord.Id
	binaryRecord.Binary = storageResponse.BinaryRecord.Binary
	binaryRecord.Name = storageResponse.BinaryRecord.Name
	binaryRecord.Note = storageResponse.BinaryRecord.Note
	binaryRecord.LastAccess = storageResponse.BinaryRecord.LastAccess
	binaryRecord.CreatedDate = storageResponse.BinaryRecord.CreatedDate

	return binaryRecord, nil
}

func (s UserData) GetBinaryRecordList(ctx context.Context) ([]pb.BinaryRecord, error) {
	var binaryRecordList []pb.BinaryRecord

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).GetBinaryRecordList(
		ctx,
		&storagePb.GetBinaryRecordListRequest{},
	)

	if err != nil {
		return nil, err
	}

	var binaryRecord pb.BinaryRecord

	for _, binaryRecordResponse := range storageResponse.BinaryRecordList {
		binaryRecord.Id = binaryRecordResponse.Id
		binaryRecord.Binary = binaryRecordResponse.Binary
		binaryRecord.Name = binaryRecordResponse.Name
		binaryRecord.Note = binaryRecordResponse.Note
		binaryRecord.LastAccess = binaryRecordResponse.LastAccess
		binaryRecord.CreatedDate = binaryRecordResponse.CreatedDate

		binaryRecordList = append(binaryRecordList, binaryRecord)
	}

	return binaryRecordList, nil
}

func (s UserData) DeleteBinaryRecordById(id string, ctx context.Context) error {
	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	_, err := (*s.storageClient.GetClient()).DeleteBinaryRecordById(
		ctx,
		&storagePb.DeleteBinaryRecordByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (s UserData) UpdateBinaryRecord(request pb.BinaryRecord, userId string, ctx context.Context) (pb.BinaryRecord, error) {
	var binaryRecord pb.BinaryRecord

	// @ToDo: find out what magic is going on here. Without a context refresh, context comes to storage service empty.
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md)
	/////////////////////////////////////////////////////////////////////////////////////////////////////

	storageResponse, err := (*s.storageClient.GetClient()).UpdateBinaryRecordById(
		ctx,
		&storagePb.UpdateBinaryRecordByIdRequest{
			CreateBinaryRecord: &storagePb.CreateBinaryRecord{
				Name:   request.Name,
				Binary: request.Binary,
				Note:   request.Note,
				UserId: userId,
			},
		},
	)

	if err != nil {
		return binaryRecord, err
	}

	binaryRecord.Id = storageResponse.BinaryRecord.Id
	binaryRecord.Name = storageResponse.BinaryRecord.Name
	binaryRecord.Binary = storageResponse.BinaryRecord.Binary
	binaryRecord.Note = storageResponse.BinaryRecord.Note
	binaryRecord.LastAccess = storageResponse.BinaryRecord.LastAccess
	binaryRecord.CreatedDate = storageResponse.BinaryRecord.CreatedDate

	return binaryRecord, nil
}
