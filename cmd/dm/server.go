package main

import (
	"errors"
	"fmt"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/aos-dev/dm/api"
	"github.com/aos-dev/dm/task"
)

// serverFlags handle flags for server command
const (
	flagHost    = "host"
	flagPort    = "port"
	flagRPCPort = "rpc-port"
)

// newServerCmd conduct server command
func newServerCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:     "server",
		Short:   fmt.Sprintf("start a http server"),
		Long:    fmt.Sprintf("dm server can start a http server to handle http request"),
		Example: "Start server: dm server",
		Args:    cobra.ExactArgs(0),
		PreRunE: func(c *cobra.Command, _ []string) error {
			return validateServerFlags()
		},
		RunE: serverRun,
	}

	serverCmd.Flags().StringP(flagHost, "h", "localhost", "server host")
	serverCmd.Flags().IntP(flagPort, "p", 7436, "web server port")
	serverCmd.Flags().Int(flagRPCPort, 7000, "grpc server port")

	// use local flags to only handle flags for current command
	serverCmd.LocalFlags().VisitAll(func(flag *pflag.Flag) {
		key := formatKeyInViper(flag.Name)
		viper.BindPFlag(key, flag)
		viper.SetDefault(key, flag.DefValue)
	})

	return serverCmd
}

func serverRun(c *cobra.Command, _ []string) error {
	logger := zapcontext.From(c.Context())

	logger.Info("start manager")

	manager, err := task.NewManager(c.Context(), task.ManagerConfig{
		Host:         viper.GetString(formatKeyInViper(flagHost)),
		GrpcPort:     viper.GetInt(formatKeyInViper(flagRPCPort)),
		DatabasePath: viper.GetString(formatKeyInViper(flagDB)),
	})
	if err != nil {
		return err
	}

	srv := api.Server{
		Host:    viper.GetString(formatKeyInViper(flagHost)),
		Port:    viper.GetInt(formatKeyInViper(flagPort)),
		DevMode: viper.GetBool(formatKeyInViper(flagDev)),
		Logger:  logger,
		DB:      manager.DB(),
		Manager: manager,
	}

	return srv.Start()
}

func validateServerFlags() error {
	db := viper.GetString(formatKeyInViper(flagDB))
	if db == "" {
		return errors.New("db flag is required")
	}
	return nil
}
