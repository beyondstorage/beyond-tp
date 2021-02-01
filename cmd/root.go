package cmd

import (
	"github.com/spf13/cobra"
)

// Name and other basic info for dm.
const (
	Name        = "dm"
	Description = "Advanced tool for data migration."
	Version     = "0.0.1"
)

var RootCmd = &cobra.Command{
	Use:     Name,
	Long:    Description,
	Version: Version,
}

func Init() error {
	RootCmd.AddCommand(ServerCmd)
	return nil
}
