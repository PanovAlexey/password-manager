package grpc

import (
	"context"
	grpcClient "user-auth/pkg/storage_grpc"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (h *UserAuthorizationHandler) Register(
	ctx context.Context,
	r *pb.RegisterRequest,
) (*pb.RegisterResponse, error) {
	var response pb.RegisterResponse

	var registerUser pb.RegisterUser
	registerUser.Email = r.RegisterUser.Email
	registerUser.Password = r.RegisterUser.Password
	registerUser.RepeatPassword = r.RegisterUser.RepeatPassword

	user.Email = "test@ya.ru"
	token, err := h.jwtAuthorizationService.GetJWTToken(user.Id)
	createUser := grpcClient.CreateUser{
		Email:          registerUser.Email,
		Password:       registerUser.Password,
		RepeatPassword: registerUser.RepeatPassword,
	}

	createUserResponse, err := (*h.storageClient.GetClient()).CreateUser(
		ctx,
		&grpcClient.CreateUserRequest{
			CreateUser: &createUser,
		},
	)

	h.logger.Info(createUserResponse)
	h.logger.Info(err)

	token, err := h.jwtAuthorizationService.GetJWTToken(createUserResponse.User.Id)

	if err != nil {
		h.logger.Error("getting JWT token by user id error: "+err.Error(), createUserResponse.User.Id)
	}

	outputUser := pb.User{
		Id:               createUserResponse.User.Id,
		Email:            createUserResponse.User.Email,
		Token:            token,
		RegistrationDate: createUserResponse.User.RegistrationDate,
		LastLogin:        createUserResponse.User.LastLogin,
	}

	response.User = &outputUser

	h.logger.Info("successful registered. ", r)

	return &response, nil
}
