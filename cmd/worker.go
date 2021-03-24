package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/aos-dev/noah/task"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// workerFlags handle flags for worker command
type workerFlags struct {
	host       string
	portalAddr string
}

var workerFlag = workerFlags{}

var WorkerCmd = &cobra.Command{
	Use:     "worker",
	Short:   fmt.Sprintf("start a task worker"),
	Long:    fmt.Sprintf("dm worker can start a task worker to handle task job distribution"),
	Example: "Start worker: dm worker",
	Args:    cobra.ExactArgs(0),
	PreRunE: func(c *cobra.Command, _ []string) error {
		return validateWorkerFlags(c)
	},
	RunE: workerRun,
}

func workerRun(c *cobra.Command, _ []string) error {
	logger := zapcontext.From(c.Context())
	logger.Info("worker info", zap.String("host", workerFlag.host),
		zap.String("portal addr", workerFlag.portalAddr))

	w, err := task.NewWorker(c.Context(), task.WorkerConfig{
		Host:       workerFlag.host,
		PortalAddr: workerFlag.portalAddr,
	})
	if err != nil {
		logger.Error("new worker failed", zap.Error(err))
		return err
	}
	err = w.Connect(c.Context())
	if err != nil {
		logger.Error("worker connect portal failed", zap.Error(err), zap.String("portal", workerFlag.portalAddr))
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

func initWorkerCmdFlags() {
	WorkerCmd.Flags().StringVarP(&workerFlag.host, "host", "h", "0.0.0.0", "worker host")
	WorkerCmd.Flags().StringVar(&workerFlag.portalAddr, "portal", "", "portal server address")
}

func validateWorkerFlags(c *cobra.Command) error {
	if portal := c.Flag("portal").Value.String(); portal == "" {
		return fmt.Errorf("portal flag is required")
	}
	return nil
}
