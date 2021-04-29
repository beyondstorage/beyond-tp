package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/task"
)

const staffCmdName = "staff"

// staffFlags handle flags for staff command
const (
	flagManager = "manager"
)

// newStaffCmd conduct staff command
func newStaffCmd() *cobra.Command {
	staffCmd := &cobra.Command{
		Use:     staffCmdName,
		Short:   fmt.Sprintf("start a task staff"),
		Long:    fmt.Sprintf("dm staff can start a task staff to handle task job distribution"),
		Example: "Start staff: dm staff",
		Args:    cobra.ExactArgs(0),
		PreRunE: func(c *cobra.Command, _ []string) error {
			return validateStaffFlags()
		},
		RunE: staffRun,
	}

	staffCmd.Flags().StringP(flagHost, "h", "localhost", "staff host")
	staffCmd.Flags().String(flagManager, "", "manager server address")

	// use local flags to only handle flags for current command
	staffCmd.LocalFlags().VisitAll(func(flag *pflag.Flag) {
		key := formatKeyInViper(staffCmdName, flag.Name)
		viper.BindPFlag(key, flag)
		viper.SetDefault(key, flag.DefValue)
	})
	return staffCmd
}

func staffRun(c *cobra.Command, _ []string) error {
	logger := zapcontext.From(c.Context())

	host, managerAddr, dbPath :=
		viper.GetString(formatKeyInViper(staffCmdName, flagHost)),
		viper.GetString(formatKeyInViper(staffCmdName, flagManager)),
		viper.GetString(formatKeyInViper("", flagDB))
	logger.Info("staff info", zap.String("host", host),
		zap.String("manager addr", managerAddr))
	w, err := task.NewStaff(c.Context(), task.StaffConfig{
		Host:        host,
		DataPath:    dbPath,
		ManagerAddr: managerAddr,
	})
	if err != nil {
		logger.Error("new staff", zap.Error(err))
		return err
	}
	err = w.Start(c.Context())
	if err != nil {
		logger.Error("staff connect manager", zap.Error(err), zap.String("manager", managerAddr))
		return err
	}

	// Setup the interrupt handler to drain so we don't miss
	// requests when scaling down.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	// TODO: We need to handle w.DisConnect here
	logger.Info("Exiting")
	return nil
}

func validateStaffFlags() error {
	if manager := viper.GetString(formatKeyInViper(staffCmdName, flagManager)); manager == "" {
		return fmt.Errorf("manager flag is required")
	}
	return dbRequiredCheck()
}
