package main

import (
	"context"
	"os"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	logger := zapcontext.From(ctx)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		logger.Error("execute command:", zap.Error(err))
		os.Exit(1)
	}
}

func init() {
	zapcontext.SetFactoryFunction(zapFactory)

	// Setup Env
	viper.SetEnvPrefix(Name)
	viper.AutomaticEnv()

	// Setup flags
	initGlobalFlags()
	initServerCmdFlags()
	initStaffCmdFlags()

	// Setup commands
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(staffCmd)
}
