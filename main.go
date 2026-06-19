package main

import (
	"os"

	"github.com/acosio14/calories-tracker-cli/cli"
)

func main() {

	if err := cli.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
