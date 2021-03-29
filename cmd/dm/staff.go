package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/aos-dev/noah/task"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// staffFlags handle flags for staff command
type staffFlags struct {
	host        string
	managerAddr string
}

var staffFlag = staffFlags{}

var staffCmd = &cobra.Command{
	Use:     "staff",
	Short:   fmt.Sprintf("start a task staff"),
	Long:    fmt.Sprintf("dm staff can start a task staff to handle task job distribution"),
	Example: "Start staff: dm staff",
	Args:    cobra.ExactArgs(0),
	PreRunE: func(c *cobra.Command, _ []string) error {
		return validateStaffFlags(c)
	},
	RunE: staffRun,
}

func staffRun(c *cobra.Command, _ []string) error {
	logger := zapcontext.From(c.Context())
	logger.Info("staff info", zap.String("host", staffFlag.host),
		zap.String("manager addr", staffFlag.managerAddr))

	w, err := task.NewStaff(c.Context(), task.StaffConfig{
		Host:        staffFlag.host,
		ManagerAddr: staffFlag.managerAddr,
	})
	if err != nil {
		logger.Error("new staff", zap.Error(err))
		return err
	}
	err = w.Connect(c.Context())
	if err != nil {
		logger.Error("staff connect manager", zap.Error(err), zap.String("manager", staffFlag.managerAddr))
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

func initStaffCmdFlags() {
	staffCmd.Flags().StringVarP(&staffFlag.host, "host", "h", "localhost", "staff host")
	staffCmd.Flags().StringVar(&staffFlag.managerAddr, "manager", "", "manager server address")
}

func validateStaffFlags(c *cobra.Command) error {
	if manager := c.Flag("manager").Value.String(); manager == "" {
		return fmt.Errorf("manager flag is required")
	}
	return nil
}
