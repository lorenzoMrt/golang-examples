package main

import (
	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "stores-cli"}
	rootCmd.AddCommand(cli.InitStoresCmd())
	rootCmd.Execute()
}
