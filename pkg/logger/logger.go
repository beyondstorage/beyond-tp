package logger

import (
	"go.uber.org/zap"
)

func init() {
	logger, _ = zap.NewProduction()
}

var logger *zap.Logger
