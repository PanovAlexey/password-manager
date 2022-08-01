package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"user-auth/internal/config"
	pb "user-auth/pkg/storage_grpc"
)

type StorageClient struct {
	client     *pb.StorageClient
	connection *grpc.ClientConn
}

func GetStorageClient(config config.Config) (StorageClient, error) {
	storageClient := StorageClient{}

	connection, err := grpc.Dial(
		config.GetStorageGRPCServerAddress(),
		//grpc.WithUnaryInterceptor( interceptors.AccessLogInterceptor @ToDo: enable interceptor ),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return storageClient, err
	}

	storageClient.connection = connection
	client := pb.NewStorageClient(connection)
	storageClient.client = &client

	return storageClient, err
}

func (c StorageClient) GetClient() *pb.StorageClient {
	return c.client
}

func (c StorageClient) GetConnection() *grpc.ClientConn {
	return c.connection
}
