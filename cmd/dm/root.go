package main

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
	db       string
	dev      bool
	logLevel string
}

var globalFlag = globalFlags{}

var rootCmd = &cobra.Command{
	Use:     Name,
	Long:    Description,
	Version: Version,
}

func initGlobalFlags() {
	rootCmd.PersistentFlags().StringVar(&globalFlag.db, "db", "", "path to locate badger db")
	rootCmd.PersistentFlags().BoolVar(&globalFlag.dev, "dev", false, "enable dev mode or not")
	rootCmd.PersistentFlags().StringVar(&globalFlag.logLevel, "log_level", "info", "log level")
	// Overwrite the default help flag to free -h shorthand.
	rootCmd.PersistentFlags().Bool("help", false, "help for this command")
}
