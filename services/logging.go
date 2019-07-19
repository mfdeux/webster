package services

import (
	"go.uber.org/zap"
)

type LoggingService interface {
}

// NewLoggingService creates a new logging service
func NewLoggingService() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	return sugar
}
