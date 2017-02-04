package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tonglil/fluxflow/versioning"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version %s\n", versioning.Get())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
