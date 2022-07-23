package logging

import (
	"api-gw/internal/config"
	"go.uber.org/zap"
	"log"
)

func GetLogger(config config.Config) *zap.SugaredLogger {
	zapLogger, err := zap.NewProduction()
	zapLogger = zapLogger.Named(config.GetApplicationName())

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	return zapLogger.Sugar()
}
