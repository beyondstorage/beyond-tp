package cmd

import (
	"github.com/spf13/cobra"

	"github.com/aos-dev/dm/constants"
)

var RootCmd = &cobra.Command{
	Use:     constants.Name,
	Long:    constants.Description,
	Version: constants.Version,
}

func Init() error {
	RootCmd.AddCommand(ServerCmd)
	return nil
}
