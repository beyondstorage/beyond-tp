package main

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Name and other basic info for dm.
const (
	Name        = "dm"
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

	rootCmd.PersistentFlags().String(flagDB, "", "path to locate badger db")
	rootCmd.PersistentFlags().Bool(flagDev, false, "enable dev mode or not")
	rootCmd.PersistentFlags().String(flagLogLevel, "info", "log level")
	// Overwrite the default help flag to free -h shorthand.
	rootCmd.PersistentFlags().Bool("help", false, "help for this command")

	rootCmd.AddCommand(newServerCmd())
	rootCmd.AddCommand(newStaffCmd())
	rootCmd.AddCommand(newTaskCmd())

	// use local flags to only handle flags for current command
	rootCmd.LocalFlags().VisitAll(func(flag *pflag.Flag) {
		key := formatKeyInViper("", flag.Name)
		viper.BindPFlag(key, flag)
		viper.SetDefault(key, flag.DefValue)
	})
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
