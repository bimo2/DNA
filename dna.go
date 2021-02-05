package main

import (
	"os"

	"github.com/bimo2/DNA/console"
	"github.com/bimo2/DNA/lib"
)

const (
	// VERSION : DNA version
	VERSION = "0.1.0"

	// BINARY : executable file name
	BINARY = "_"
)

func main() {
	dnaFile := lib.Find()
	argv := os.Args[1:]

	if len(argv) < 1 {
		return
	}

	switch argv[0] {
	case "version":
		console.Message("version " + VERSION + " (MIT)")

	case "init":
		if dnaFile == nil {
			lib.Create()
		}

		console.Message("dna.json configured")

	default:
		return
	}
}
