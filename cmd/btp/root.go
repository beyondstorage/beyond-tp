package main

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Name and other basic info for btp.
const (
	Name        = "btp"
	Description = "Advanced tool for data migration."
	Version     = "0.2.0"
)

// globalFlags handle flags for global command
const (
	flagDB       = "db"
	flagDev      = "dev"
	flagLogLevel = "log-level"
)

// newRootCmd conduct rootCmd
func newRootCmd() *cobra.Command {
	// Setup Env
	viper.SetEnvPrefix(Name)
	viper.AutomaticEnv()

	rootCmd := &cobra.Command{
		Use:     Name,
		Long:    Description,
		Version: Version,
	}

	rootCmd.AddCommand(newAgentCmd())
	rootCmd.AddCommand(newTaskCmd())
	return rootCmd
}

// dbRequiredCheck check db flag
func dbRequiredCheck() error {
	db := viper.GetString(formatKeyInViper("", flagDB))
	if db == "" {
		return errors.New("db flag is required")
	}
	return nil
}
