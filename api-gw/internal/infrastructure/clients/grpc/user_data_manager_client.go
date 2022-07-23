package grpc

import (
	"api-gw/internal/config"
	pb "api-gw/pkg/user_data_manager_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserDataManagerClient struct {
	client     *pb.UserDataManagerClient
	connection *grpc.ClientConn
}

func GetUserDataManagerClient(config config.Config) (UserDataManagerClient, error) {
	userDataManagerClient := UserDataManagerClient{}

	connection, err := grpc.Dial(
		config.GetUserDataManagerGRPCServerAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return userDataManagerClient, err
	}

	userDataManagerClient.connection = connection
	client := pb.NewUserDataManagerClient(connection)
	userDataManagerClient.client = &client

	return userDataManagerClient, err
}

func (c UserDataManagerClient) GetClient() *pb.UserDataManagerClient {
	return c.client
}

func (c UserDataManagerClient) GetConnection() *grpc.ClientConn {
	return c.connection
}
