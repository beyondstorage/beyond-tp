package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const loggerGinCtxKey = "logger_in_gin"

// MustGetLoggerFromGin wrap the MustGet method and type assert
func MustGetLoggerFromGin(c *gin.Context) *zap.Logger {
	return c.MustGet(loggerGinCtxKey).(*zap.Logger)
}

// SetLoggerInGin set a zap logger into context
func SetLoggerInGin(c *gin.Context, logger *zap.Logger) {
	c.Set(loggerGinCtxKey, logger)
}
