package main

import (
	"github.com/tonglil/fluxflow/cmd"
	"github.com/tonglil/fluxflow/versioning"
)

// Deliberately uninitialized, see versioning package.
var version string

func main() {
	versioning.Set(version)

	cmd.Execute()
}
