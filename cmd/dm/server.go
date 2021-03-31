package main

import (
	"errors"
	"fmt"

	"github.com/aos-dev/dm/task"
	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/api"
	"github.com/aos-dev/dm/models"
)

// serverFlags handle flags for server command
type serverFlags struct {
	host      string
	port      int
	rpcPort   int
	queuePort int
}

var serverFlag = serverFlags{}

var serverCmd = &cobra.Command{
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

func serverRun(c *cobra.Command, _ []string) error {
	logger := zapcontext.From(c.Context())

	db, err := models.NewDB(globalFlag.db)
	if err != nil {
		logger.Error("new db:", zap.Error(err), zap.String("path", globalFlag.db))
		return err
	}

	logger.Info("start manager")
	manager, err := task.NewManager(c.Context(), task.ManagerConfig{
		Host:      serverFlag.host,
		GrpcPort:  serverFlag.rpcPort,
		QueuePort: serverFlag.queuePort,
	})
	if err != nil {
		return err
	}

	srv := api.Server{
		Host:    serverFlag.host,
		Port:    serverFlag.port,
		DevMode: globalFlag.dev,
		Logger:  logger,
		DB:      db,
		Manager: manager,
	}

	return srv.Start()
}

func initServerCmdFlags() {
	serverCmd.Flags().StringVarP(&serverFlag.host, "host", "h", "localhost", "server host")
	serverCmd.Flags().IntVarP(&serverFlag.port, "port", "p", 7436, "web server port")
	serverCmd.Flags().IntVar(&serverFlag.rpcPort, "rpc-port", 7000, "grpc server port")
	serverCmd.Flags().IntVar(&serverFlag.queuePort, "queue-port", 7010, "msg queue server port")
}

func validateServerFlags(c *cobra.Command) error {
	if db := c.Flag("db").Value.String(); db == "" {
		return errors.New("db flag is required")
	}
	return nil
}
