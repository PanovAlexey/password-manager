package logging

import (
	"api-gw/internal/config"
	"encoding/json"
	"go.uber.org/zap"
	"log"
)

func GetLogger(config config.Config) *zap.SugaredLogger {
	rawJSON := []byte(`{
   "level": "debug",
   "encoding": "json",
   "outputPaths": ["stdout"],
   "errorOutputPaths": ["stderr"],
   "encoderConfig": {
     "messageKey": "message",
     "levelKey": "level",
     "levelEncoder": "lowercase"
   }
 }`)
	var cfg zap.Config

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	zapLogger, err := cfg.Build()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	zapLogger = zapLogger.Named(config.GetApplicationName())

	return zapLogger.Sugar()
}
