package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/beyondstorage/beyond-tp/http"
	"github.com/beyondstorage/beyond-tp/rpc"
	"github.com/beyondstorage/beyond-tp/task"
)

const (
	agentFlagRole     = "role"
	agentFlagRpcAddr  = "rpc-addr"
	agentFlagHttpAddr = "http-addr"
	agentFlagDataDir  = "data-dir"
)

var agentFlags = []cli.Flag{
	&cli.StringSliceFlag{
		Name:    agentFlagRole,
		Usage:   "role of this agent, available value: server and client",
		EnvVars: []string{"BTP_AGENT_ROLE"},
		Value:   cli.NewStringSlice("server", "client"),
	},
	&cli.StringFlag{
		Name:    agentFlagRpcAddr,
		Usage:   "rpc addr used for internal RPC communication server and client ",
		EnvVars: []string{"BTP_AGENT_RPC_ADDR"},
		Value:   "localhost:4100",
	},
	&cli.StringFlag{
		Name:    agentFlagHttpAddr,
		Usage:   "http addr used for web console and graphql API.",
		EnvVars: []string{"BTP_AGENT_HTTP_ADDR"},
		Value:   "localhost:4000",
	},
	&cli.StringFlag{
		Name:    agentFlagDataDir,
		Usage:   "data dir used storing agent's runtime data.",
		EnvVars: []string{"BTP_AGENT_DATA_DIR"},
		Value:   filepath.Join(userConfigDir(), "btp", "data"),
	},
}

var agentCmd = &cli.Command{
	Name:  "agent",
	Usage: "btp agent",
	Flags: agentFlags,
	Action: func(c *cli.Context) error {
		logger, _ := zap.NewDevelopment()

		role := parseRole(c.StringSlice(agentFlagRole))

		// Serve a server role agent.
		if _, ok := role["server"]; ok {
			err := setupServer(c, logger)
			if err != nil {
				return fmt.Errorf("setup agent server: %w", err)
			}
		}

		// Serve a client role agent.
		if _, ok := role["client"]; ok {
			err := setupClient(c, logger)
			if err != nil {
				return fmt.Errorf("setup agent client: %w", err)
			}
		}

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, os.Kill)
		<-quit
		logger.Warn("server shutdown...")
		return nil
	},
}

func setupServer(c *cli.Context, logger *zap.Logger) error {
	ts, err := task.NewServer(&task.ServerConfig{
		DataDir: c.String(agentFlagDataDir),
		Logger:  logger,
	})
	if err != nil {
		return fmt.Errorf("task new server: %w", err)
	}

	rpcSrv := rpc.New(&rpc.Config{
		Addr:   c.String(agentFlagRpcAddr),
		Task:   ts,
		Logger: logger,
	})
	go func() {
		err = rpcSrv.Serve(context.Background())
		if err != nil {
			logger.Error("rpc serve", zap.Error(err))
		}
	}()

	httpSrv := http.New(&http.Config{
		Addr:   c.String(agentFlagHttpAddr),
		Task:   ts,
		Logger: logger,
	})
	go func() {
		err = httpSrv.Serve()
		if err != nil {
			logger.Error("http serve", zap.Error(err))
		}
	}()
	return nil
}

func setupClient(c *cli.Context, logger *zap.Logger) error {
	tc, err := task.NewClient(&task.ClientConfig{
		Addr:   c.String(agentFlagRpcAddr),
		Logger: logger,
	})
	if err != nil {
		return fmt.Errorf("task new client: %w", err)
	}
	go func() {
		err = tc.Start(context.Background())
		if err != nil {
			logger.Error("client start", zap.Error(err))
		}
	}()
	return nil
}

func parseRole(s []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}
