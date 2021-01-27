package main

import (
	"context"
	"os"

	"github.com/aos-dev/dm/cmd"
)

func main() {
	if err := cmd.Init(); err != nil {
		println(err)
		os.Exit(1)
	}

	err := cmd.RootCmd.ExecuteContext(context.Background())
	if err != nil {
		os.Exit(1)
	}
}
