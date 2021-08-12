package logger

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, args ...zap.Field) {
	logger.Info(message, args...)
}

func Warn(message string, args ...zap.Field) {
	logger.Warn(message, args...)
}

func Debug(message string, args ...zap.Field) {
	logger.Debug(message, args...)
}

func Error(message string, args ...zap.Field) {
	logger.Error(message, args...)
}
