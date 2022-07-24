package servers

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"user-data-manager/internal/config"
	grpcHandler "user-data-manager/internal/handlers/grpc"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

func RunGrpcServer(config config.Config, logger Logger) {
	logger.Info(config.GetApplicationName() + " grpc server starting...")
	listen, err := net.Listen("tcp", config.GetGrpcServerAddress())

	if err != nil {
		logger.Error("grpc listening error: " + err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterUserDataManagerServer(s, grpcHandler.GetUserDataManagerHandler(logger))

	logger.Info(config.GetApplicationName() + " grpc server started")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.Serve(listen); err != nil {
			logger.Error("grpc server error: " + err.Error())
		}
	}()

	logger.Info("Signal detected: ", <-sigs)
	logger.Info(config.GetApplicationName() + " grpc server is shutdowning...")
	s.GracefulStop()
	logger.Info(config.GetApplicationName() + " grpc server has been stopped.")
}
