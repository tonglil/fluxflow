package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Output shell completion code for tab completion",
	Long: `Print shell code for evaluation of interactive completion of fluxflow commands

Examples:
  $ source <(fluxflow completion)

  Note that this depends on the bash-completion framework. It must be sourced
  before sourcing the fluxflow completion, e.g. on the Mac:

  $ brew install bash-completion
  $ source $(brew --prefix)/etc/bash_completion
  $ source <(fluxflow completion bash)`,
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenBashCompletion(os.Stdout)
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
