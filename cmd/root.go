package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:           "fluxflow",
	Short:         "A command line tool to adjust f.lux settings",
	SilenceErrors: true,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
