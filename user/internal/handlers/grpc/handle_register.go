package grpc

import (
	"context"
	"errors"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (h *UserAuthorizationHandler) Register(
	ctx context.Context,
	r *pb.RegisterRequest,
) (*pb.RegisterResponse, error) {
	var registerUser pb.RegisterUser
	registerUser.Email = r.RegisterUser.Email
	registerUser.Password = r.RegisterUser.Password
	registerUser.RepeatPassword = r.RegisterUser.RepeatPassword

	err := h.userRegistrationService.Validate(registerUser)

	if err != nil {
		return nil, err
	}

	createdUser, err := h.userRegistrationService.Register(registerUser, ctx)

	if err != nil {
		return nil, err
	}

	token, err := h.jwtAuthorizationService.GetJWTToken(createdUser.Id)

	if err != nil {
		h.logger.Error("getting JWT token by user id error: "+err.Error(), createdUser.Id)
		return nil, errors.New("getting JWT token by user id error: " + err.Error())
	}

	outputUser := pb.User{
		Id:               createdUser.Id,
		Email:            createdUser.Email,
		Token:            token,
		RegistrationDate: createdUser.RegistrationDate,
		LastLogin:        createdUser.LastLogin,
	}

	h.logger.Info("successful registered. ", outputUser)

	return &pb.RegisterResponse{User: &outputUser}, nil
}
