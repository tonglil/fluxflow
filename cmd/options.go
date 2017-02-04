package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tonglil/fluxflow/menu"
	"github.com/tonglil/fluxflow/menu/submenu"
	"github.com/tonglil/fluxflow/pkg"
)

func init() {
	RootCmd.AddCommand(hourCmd)
	RootCmd.AddCommand(sunriseCmd)
	RootCmd.AddCommand(movieCmd)
	RootCmd.AddCommand(darkroomCmd)
}

func success(cmd *cobra.Command, args []string) {
	fmt.Println("Updated f.lux settings.")
}

var hourCmd = &cobra.Command{
	Use:   "hour",
	Short: "Disable f.lux for 1 hour",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.Run(menu.Disable, submenu.Hour)
	},
	PostRun: success,
}

var sunriseCmd = &cobra.Command{
	Use:   "sunrise",
	Short: "Disable f.lux until sunrise",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.Run(menu.Disable, submenu.Sunrise)
	},
	PostRun: success,
}

var movieCmd = &cobra.Command{
	Use:   "movie",
	Short: "Set f.lux to movie mode for 2.5 hours",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.Run(menu.Color, submenu.Movie)
	},
	PostRun: success,
}

var darkroomCmd = &cobra.Command{
	Use:   "darkroom",
	Short: "Set f.lux to darkroom mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.Run(menu.Color, submenu.Darkroom)
	},
	PostRun: success,
}
