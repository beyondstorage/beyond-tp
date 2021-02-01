package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aos-dev/dm/api"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Short:   fmt.Sprintf("start a http server"),
	Long:    fmt.Sprintf("dm server can start a http server to handle http request"),
	Example: "Start server: dm server",
	Args:    cobra.ExactArgs(0),
	RunE:    serverRun,
}

func serverRun(_ *cobra.Command, _ []string) error {
	return api.StartServer()
}
