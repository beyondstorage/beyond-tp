package main

import (
	"fmt"

	"github.com/beyondstorage/go-toolbox/zapcontext"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/beyondstorage/beyond-tp/models"
)

const (
	taskCmdName     = "task"
	taskInfoCmdName = "info"
	taskListCmdName = "list"
	taskRunCmdName  = "run"
)

// newTaskCmd conduct task command
func newTaskCmd() *cobra.Command {
	taskCmd := &cobra.Command{
		Use:  taskCmdName,
		Long: fmt.Sprintf("Sub commands about task"),
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			return dbRequiredCheck()
		},
	}

	taskCmd.AddCommand(newTaskInfoCmd())
	taskCmd.AddCommand(newTaskListCmd())
	taskCmd.AddCommand(newTaskRunCmd())
	return taskCmd
}

func newTaskListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     taskListCmdName,
		Short:   fmt.Sprintf("list tasks"),
		Long:    fmt.Sprintf("beyondtp task list can list all task info"),
		Example: "List tasks: beyondtp task list",
		Args:    cobra.ExactArgs(0),
		RunE:    taskListRun,
	}
	return listCmd
}

func taskListRun(c *cobra.Command, _ []string) error {
	logger := zapcontext.From(c.Context())

	dbPath := viper.GetString(formatKeyInViper("", flagDB))
	db, err := models.NewDB(dbPath, logger)
	if err != nil {
		logger.Error("create db", zap.String("path", dbPath), zap.Error(err))
		return err
	}

	tasks, err := db.ListTasks()
	if err != nil {
		logger.Error("list task", zap.Error(err))
		return err
	}

	for _, t := range tasks {
		fmt.Fprintf(c.OutOrStdout(), "%s\n", t.String())
	}

	return nil
}

func newTaskInfoCmd() *cobra.Command {
	infoCmd := &cobra.Command{
		Use:     taskInfoCmdName,
		Short:   fmt.Sprintf("get task info by ID"),
		Long:    fmt.Sprintf("beyondtp task info can get a task's info by task ID"),
		Example: "Show task info: beyondtp task info [task-ID] [task-ID]...",
		Args:    cobra.MinimumNArgs(1),
		RunE:    taskInfoRun,
	}
	return infoCmd
}

func taskInfoRun(c *cobra.Command, args []string) error {
	logger := zapcontext.From(c.Context())

	dbPath := viper.GetString(formatKeyInViper("", flagDB))
	db, err := models.NewDB(dbPath, logger)
	if err != nil {
		logger.Error("create db", zap.String("path", dbPath), zap.Error(err))
		return err
	}

	for _, id := range args {
		task, err := db.GetTask(id)
		if err != nil {
			logger.Error("get task", zap.String("id", id), zap.Error(err))
			continue
		}
		fmt.Fprintf(c.OutOrStdout(), "%s\n", task.String())
	}

	return nil
}

func newTaskRunCmd() *cobra.Command {
	runCmd := &cobra.Command{
		Use:     taskRunCmdName,
		Short:   fmt.Sprintf("run task by ID"),
		Long:    fmt.Sprintf("beyondtp task run can run a task by task ID"),
		Example: "Run task: beyondtp task run [task-ID]",
		Args:    cobra.ExactArgs(1),
		RunE:    taskRunRun,
	}
	return runCmd
}

func taskRunRun(c *cobra.Command, args []string) error {
	logger := zapcontext.From(c.Context())

	dbPath := viper.GetString(formatKeyInViper("", flagDB))
	db, err := models.NewDB(dbPath, logger)
	if err != nil {
		logger.Error("create db", zap.String("path", dbPath), zap.Error(err))
		return err
	}

	taskID := args[0]
	task, err := db.GetTask(taskID)
	if err != nil {
		logger.Error("get task", zap.String("id", taskID), zap.Error(err))
		return err
	}

	if task.Status != models.TaskStatus_Created {
		err = fmt.Errorf("task status Created expected, but got <%s>", task.Status.String())
		logger.Error("task status check", zap.String("id", taskID), zap.Error(err))
		return err
	}

	task.UpdatedAt = timestamppb.Now()
	task.Status = models.TaskStatus_Ready

	err = db.UpdateTask(nil, task)
	if err != nil {
		logger.Error("save task", zap.String("id", task.Id), zap.Error(err))
		return err
	}

	fmt.Fprintf(c.OutOrStdout(), "task <%s> started\n", taskID)
	return nil
}
