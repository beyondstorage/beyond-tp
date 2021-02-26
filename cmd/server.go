package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aos-dev/dm/api"
	ilog "github.com/aos-dev/dm/pkg/logger"
)

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
	RunE:    serverRun,
}

func serverRun(c *cobra.Command, _ []string) error {
	cfg := api.ServerConfig{
		Host:   serverFlag.host,
		Port:   serverFlag.port,
		Logger: ilog.FromContext(c.Context()),
	}

	return api.StartServer(cfg)
}

func initServerCmdFlags() {
	ServerCmd.Flags().StringVarP(&serverFlag.host, "host", "h", "0.0.0.0", "server host")
	ServerCmd.Flags().IntVarP(&serverFlag.port, "port", "p", 7436, "server port")
}
