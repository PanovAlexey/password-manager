package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	grpcServer "user-data-manager/internal/servers/grpc"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserDataManagerServer(s, grpcServer.GetUserDataManagerServer())

	fmt.Println("Server gRPC started...")

	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
