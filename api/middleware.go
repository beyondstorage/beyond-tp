package api

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/beyondstorage/go-toolbox/zapcontext"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const requestIDCtxKey = "request_id_ctx_key"

// setRequestID init requestID and set into context
func setRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.NewString()
		c.Writer.Header().Set("x-dm-request-id", requestID) // set uuid in header for frontend use
		c.Set(requestIDCtxKey, requestID)
		c.Next()
	}
}

// setLoggerWithReqID set a zap logger with requestID into context
func setLoggerWithReqID(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		loggerWithRequestID := logger.With(zap.String("request_id", c.GetString(requestIDCtxKey)))
		zapcontext.WithinGin(c, loggerWithRequestID)
		c.Next()
	}
}

// logRequestInfo inspired from `https://github.com/gin-contrib/zap`
func logRequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger := zapcontext.FromGin(c)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// ginRecovery inspired from `https://github.com/gin-contrib/zap`
func ginRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				logger := zapcontext.FromGin(c)
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				logger.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
