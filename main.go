package main

import (
	"os"

	"github.com/simioa/scriptbeat/cmd"

	_ "github.com/simioa/scriptbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
