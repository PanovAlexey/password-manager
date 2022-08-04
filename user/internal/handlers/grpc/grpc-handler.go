package grpc

import (
	"user-auth/internal/application/service"
	pb "user-auth/pkg/user_authorization_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type UserAuthorizationHandler struct {
	logger                               Logger
	jwtAuthorizationService              service.JWTAuthorization
	userRegistrationService              service.UserRegistration
	userMetadataFromContextGetterService service.UserMetadataFromContextGetter
	pb.UnimplementedUserAuthorizationServer
}

func GetUserAuthorizationHandler(
	logger Logger,
	jwtAuthorizationService service.JWTAuthorization,
	userRegistrationService service.UserRegistration,
	userMetadataFromContextGetterService service.UserMetadataFromContextGetter,
) *UserAuthorizationHandler {
	return &UserAuthorizationHandler{
		logger:                               logger,
		jwtAuthorizationService:              jwtAuthorizationService,
		userRegistrationService:              userRegistrationService,
		userMetadataFromContextGetterService: userMetadataFromContextGetterService,
	}
}
