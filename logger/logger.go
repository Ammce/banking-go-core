package logger

import "go.uber.org/zap"

var log *zap.Logger

func init() {
	var err error
	log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, args ...zap.Field) {
	log.Info(message, args...)
}

func Warn(message string, args ...zap.Field) {
	log.Warn(message, args...)
}

func Debug(message string, args ...zap.Field) {
	log.Debug(message, args...)
}

func Error(message string, args ...zap.Field) {
	log.Error(message, args...)
}
