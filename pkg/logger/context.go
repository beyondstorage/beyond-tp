package logger

import (
	"context"

	"go.uber.org/zap"
)

type contextKey struct{}

// loggerKey is used as key to store logger in context
var loggerKey contextKey

// WithinContext set *Logger into given context and return
func WithinContext(ctx context.Context, l *zap.Logger) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	// if nil logger was given, return ctx directly
	if l == nil {
		return ctx
	}

	return context.WithValue(ctx, loggerKey, l)
}

// FromContext get *Logger from context
// Notice: If ctx is nil or no Logger was set before, it will return a default logger
func FromContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	l, ok := ctx.Value(loggerKey).(*zap.Logger)
	if !ok {
		return logger
	}
	return l
}
