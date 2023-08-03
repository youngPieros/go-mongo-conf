package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logger = func() func() *zap.SugaredLogger {
	var logger *zap.Logger
	return func() *zap.SugaredLogger {
		if logger == nil {
			loggerConfig := zap.NewProductionConfig()
			loggerConfig.EncoderConfig.TimeKey = "timestamp"
			loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
			logger, _ = loggerConfig.Build()
		}
		return logger.Sugar()
	}
}()()
