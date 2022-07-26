package grpc

import (
	"api-gw/internal/config"
	pb "api-gw/pkg/user_authorization_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserAuthorizationClient struct {
	client     *pb.UserAuthorizationClient
	connection *grpc.ClientConn
}

func GetUserAuthorizationClient(config config.Config) (UserAuthorizationClient, error) {
	userAuthorizationClient := UserAuthorizationClient{}

	connection, err := grpc.Dial(
		config.GetUserAuthorizationGRPCServerAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return userAuthorizationClient, err
	}

	userAuthorizationClient.connection = connection
	client := pb.NewUserAuthorizationClient(connection)
	userAuthorizationClient.client = &client

	return userAuthorizationClient, err
}

func (c UserAuthorizationClient) GetClient() *pb.UserAuthorizationClient {
	return c.client
}

func (c UserAuthorizationClient) GetConnection() *grpc.ClientConn {
	return c.connection
}
