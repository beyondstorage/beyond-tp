package main

import (
	"context"
	"os"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"go.uber.org/zap"
)

func main() {
	zapcontext.SetFactoryFunction(zapFactory)

	ctx := context.Background()
	logger := zapcontext.From(ctx)

	if err := newRootCmd().ExecuteContext(ctx); err != nil {
		logger.Error("execute command:", zap.Error(err))
		os.Exit(1)
	}
}
