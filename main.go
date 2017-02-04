package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/tonglil/fluxflow/menu"
	"github.com/tonglil/fluxflow/menu/submenu"
	"github.com/tonglil/fluxflow/script"
)

func main() {
	primary := menu.Color
	secondary := submenu.Darkroom

	cmd := exec.Command("osascript", "-e", fmt.Sprintf(script.Flux, primary, secondary))
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated Flux settings.")
}
