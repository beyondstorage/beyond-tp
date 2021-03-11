package cmd

import (
	"errors"
	"fmt"
	"github.com/aos-dev/dm/models"
	"go.uber.org/zap"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/spf13/cobra"

	"github.com/aos-dev/dm/api"
)

// serverFlags handle flags for server command
type serverFlags struct {
	host string
	port int
}

var serverFlag = serverFlags{}

var ServerCmd = &cobra.Command{
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
	cfg := api.Config{
		Host:   serverFlag.host,
		Port:   serverFlag.port,
		Debug:  globalFlag.debug,
		DBPath: globalFlag.db,
	}

	ctx := c.Context()
	logger := zapcontext.From(ctx)

	db, err := models.NewDB(cfg.DBPath)
	if err != nil {
		logger.Fatal("new db failed:", zap.Error(err), zap.String("path", cfg.DBPath))
	}

	ctx = models.DbIntoContext(ctx, db)

	return api.StartServer(ctx, cfg)
}

func initServerCmdFlags() {
	ServerCmd.Flags().StringVarP(&serverFlag.host, "host", "h", "0.0.0.0", "server host")
	ServerCmd.Flags().IntVarP(&serverFlag.port, "port", "p", 7436, "server port")
}

func validateServerFlags(c *cobra.Command) error {
	if db := c.Flag("db").Value.String(); db == "" {
		return errors.New("db flag is required")
	}
	return nil
}
