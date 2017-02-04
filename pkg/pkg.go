package pkg

import (
	"fmt"
	"os/exec"

	"github.com/tonglil/fluxflow/menu"
	"github.com/tonglil/fluxflow/menu/submenu"
	"github.com/tonglil/fluxflow/script"
)

// Run executes the AppleScript with the chosen options.
func Run(primary menu.Item, secondary submenu.Item) error {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(script.Flux, primary, secondary))

	return cmd.Run()
}
