package main

import (
	"os"

	"github.com/yasvanth/shibidpbeat/cmd"

	_ "github.com/yasvanth/shibidpbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
