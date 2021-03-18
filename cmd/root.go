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

// globalFlags handle flags for global command
type globalFlags struct {
	db    string
	debug bool
}

var globalFlag = globalFlags{}

var RootCmd = &cobra.Command{
	Use:     Name,
	Long:    Description,
	Version: Version,
}

func Init() error {
	initGlobalFlags()
	initServerCmdFlags()
	initWorkerCmdFlags()

	RootCmd.AddCommand(ServerCmd)
	RootCmd.AddCommand(WorkerCmd)
	return nil
}

func initGlobalFlags() {
	RootCmd.PersistentFlags().StringVar(&globalFlag.db, "db", "", "path to locate badger db")
	RootCmd.PersistentFlags().BoolVar(&globalFlag.debug, "debug", false, "enable debug or not")
	// Overwrite the default help flag to free -h shorthand.
	RootCmd.PersistentFlags().Bool("help", false, "help for this command")
}
