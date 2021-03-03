package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const loggerGinCtxKey = "logger_in_gin"

// FromGinContext wrap the MustGet method and type assert
func FromGinContext(c *gin.Context) *zap.Logger {
	return c.MustGet(loggerGinCtxKey).(*zap.Logger)
}

// WithinGinContext set a zap logger into context
func WithinGinContext(c *gin.Context, logger *zap.Logger) {
	c.Set(loggerGinCtxKey, logger)
}
