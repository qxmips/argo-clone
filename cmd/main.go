package main

import (
	"os"

	apiserver "github.com/qxmips/cmd/argocd-server/commands"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
)

func main() {
	var command *cobra.Command
	command = apiserver.NewCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
