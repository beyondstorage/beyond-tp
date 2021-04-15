package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Name and other basic info for dm.
const (
	Name        = "dm"
	Description = "Advanced tool for data migration."
	Version     = "0.2.0"
)

// globalFlags handle flags for global command
type globalFlags struct {
	db       string
	dev      bool
	logLevel string
}

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

	rootCmd.PersistentFlags().String("db", "", "path to locate badger db")
	rootCmd.PersistentFlags().Bool("dev", false, "enable dev mode or not")
	rootCmd.PersistentFlags().String("log-level", "info", "log level")
	// Overwrite the default help flag to free -h shorthand.
	rootCmd.PersistentFlags().Bool("help", false, "help for this command")

	// bind log-level flag with env key log_level (with prefix dm)
	viper.BindPFlag("log_level", rootCmd.Flag("log-level"))

	rootCmd.AddCommand(newServerCmd())
	rootCmd.AddCommand(newStaffCmd())
	return rootCmd
}

// parseGlobalFlag get flag values from command flags
func parseGlobalFlag(c *cobra.Command) (globalFlags, error) {
	flags := c.Flags()
	db, err := flags.GetString("db")
	if err != nil {
		return globalFlags{}, err
	}

	dev, err := flags.GetBool("dev")
	if err != nil {
		return globalFlags{}, err
	}

	logLevel, err := flags.GetString("log-level")
	if err != nil {
		return globalFlags{}, err
	}

	return globalFlags{
		db:       db,
		dev:      dev,
		logLevel: logLevel,
	}, nil
}
