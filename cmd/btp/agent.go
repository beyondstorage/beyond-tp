package main

import (
	"github.com/spf13/cobra"
)

func newAgentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "agent",
		Short:   "Start an agent",
		Example: "Start agent: btp agent",
		Args:    cobra.ExactArgs(0),
		RunE:    agentRun,
	}

	// Agent related config
	cmd.Flags().String("config", "", "The config path for agent")
	cmd.Flags().String("log_level", "", "Log level for agent")
	cmd.Flags().Bool("dev", false, "Enable dev mode or not")
	cmd.Flags().String("data_dir", "", "The data dir for agent")

	// Client related config
	cmd.Flags().Bool("client", false, "Only enable client role")

	// Server related config
	cmd.Flags().Bool("server", false, "Only enable server role")
	cmd.Flags().String("ui_addr", "localhost:7000", "The UI addr for server")
	cmd.Flags().String("rpc_addr", "localhost:7100", "The RPC addr for server")

	return cmd
}

func agentRun(c *cobra.Command, _ []string) error {
	return nil
}
