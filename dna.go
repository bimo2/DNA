package main

import (
	"os"

	"github.com/bimo2/DNA/cli"
	"github.com/bimo2/DNA/console"
)

const (
	// VERSION : DNA version
	VERSION = "0.1.0"

	// FILENAME : DNA config filename
	FILENAME = "dna.json"

	// BINARY : executable file name
	BINARY = "_"
)

func main() {
	dnaFile := cli.Load(FILENAME)
	argv := os.Args[1:]

	if len(argv) < 1 {
		return
	}

	switch argv[0] {
	case "version":
		console.Message("version " + VERSION + " (MIT)")

	case "init":
		if dnaFile == nil {
			cli.Initialize(FILENAME)
		}

		console.Message("\"" + FILENAME + "\" configured")

	default:
		return
	}
}
