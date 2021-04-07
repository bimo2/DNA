package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bimo2/DNA/cli"
	"github.com/bimo2/DNA/console"
	"github.com/bimo2/DNA/protocol"
)

const (
	// VERSION : DNA version
	VERSION = "0.2.1"

	// FILENAME : DNA config filename
	FILENAME = "dna.json"

	// BINARY : executable file name
	BINARY = "_"
)

func main() {
	config, path, err := protocol.Find(FILENAME)

	if err != nil {
		console.Error(err.Error())
		return
	}

	argv := os.Args[1:]

	if len(argv) < 1 {
		listCommands(config != nil)
		return
	}

	switch argv[0] {
	case "init", "i":
		if config == nil {
			cli.Init(FILENAME)
		} else {
			message := "`" + FILENAME + "` already exists"
			console.Error(message)
		}

	case "list", "ls":
		if notFound(config) {
			return
		}

		cli.List(&config.Scripts)

	case "version", "v":
		version := "version " + VERSION + " (" + runtime.GOOS + ")"
		console.Message(version, nil)

	default:
		if notFound(config) {
			return
		}

		name := argv[0]
		script, exists := config.Scripts[name]

		if !exists {
			console.Error("`" + name + "` not defined")
			return
		}

		cli.ExecSync(&argv, &script, &config.Env, path)
	}
}

func listCommands(init bool) {
	if init {
		console.Message("Configured!", nil)
	} else {
		console.Message("Not configured", nil)
		fmt.Println("\n- " + console.BOLD + "init, i\n  " + console.DEFAULT + "Create `dna.json` template")
	}

	fmt.Println("\n- " + console.BOLD + "list, ls\n  " + console.DEFAULT + "List all project scripts")
	fmt.Println("\n- " + console.BOLD + "version, v\n  " + console.DEFAULT + "Get version information")
	fmt.Println()
}

func notFound(config *protocol.DNAFile) bool {
	if config == nil {
		console.Error("`" + FILENAME + "` not found")
		return true
	}

	return false
}
