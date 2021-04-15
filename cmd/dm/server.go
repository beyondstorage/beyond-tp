package main

import (
	"errors"
	"fmt"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/api"
	"github.com/aos-dev/dm/task"
)

// serverFlags handle flags for server command
type serverFlags struct {
	host      string
	port      int
	rpcPort   int
	queuePort int
}

// newServerCmd conduct server command
func newServerCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:     "server",
		Short:   fmt.Sprintf("start a http server"),
		Long:    fmt.Sprintf("dm server can start a http server to handle http request"),
		Example: "Start server: dm server",
		Args:    cobra.ExactArgs(0),
		PreRunE: func(c *cobra.Command, _ []string) error {
			return validateServerFlags(c)
		},
		RunE: serverRun,
	}

	serverCmd.Flags().StringP("host", "h", "localhost", "server host")
	serverCmd.Flags().IntP("port", "p", 7436, "web server port")
	serverCmd.Flags().Int("rpc-port", 7000, "grpc server port")
	serverCmd.Flags().Int("queue-port", 7010, "msg queue server port")

	return serverCmd
}

func serverRun(c *cobra.Command, _ []string) error {
	logger := zapcontext.From(c.Context())

	logger.Info("start manager")

	serverFlag, err := parseServerFlags(c)
	if err != nil {
		logger.Error("parse flag for server command", zap.Error(err))
		return err
	}

	globalFlag, err := parseGlobalFlag(c)
	if err != nil {
		logger.Error("parse global flag", zap.Error(err))
		return err
	}

	manager, err := task.NewManager(c.Context(), task.ManagerConfig{
		Host:         serverFlag.host,
		GrpcPort:     serverFlag.rpcPort,
		DatabasePath: globalFlag.db,
	})
	if err != nil {
		return err
	}

	srv := api.Server{
		Host:    serverFlag.host,
		Port:    serverFlag.port,
		DevMode: globalFlag.dev,
		Logger:  logger,
		DB:      manager.DB(),
		Manager: manager,
	}

	return srv.Start()
}

func validateServerFlags(c *cobra.Command) error {
	if db := c.Flag("db").Value.String(); db == "" {
		return errors.New("db flag is required")
	}
	return nil
}

// parseServerFlags get flag values from command flags
func parseServerFlags(c *cobra.Command) (serverFlags, error) {
	flagSet := c.Flags()
	host, err := flagSet.GetString("host")
	if err != nil {
		return serverFlags{}, err
	}
	port, err := flagSet.GetInt("port")
	if err != nil {
		return serverFlags{}, err
	}
	rpcPort, err := flagSet.GetInt("rpc-port")
	if err != nil {
		return serverFlags{}, err
	}
	queuePort, err := flagSet.GetInt("queue-port")
	if err != nil {
		return serverFlags{}, err
	}
	return serverFlags{
		host:      host,
		port:      port,
		rpcPort:   rpcPort,
		queuePort: queuePort,
	}, nil
}
