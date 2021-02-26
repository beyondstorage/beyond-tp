package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const loggerCtxKey = "logger_in_ctx"

// mustGetLoggerFrom wrap the MustGet method and type assert
func mustGetLoggerFrom(c *gin.Context) *zap.Logger {
	return c.MustGet(loggerCtxKey).(*zap.Logger)
}

// setLoggerIn set a zap logger into context
func setLoggerIn(c *gin.Context, logger *zap.Logger) {
	c.Set(loggerCtxKey, logger)
}
