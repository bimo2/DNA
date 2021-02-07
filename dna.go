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
	VERSION = "0.1.0"

	// FILENAME : DNA config filename
	FILENAME = "dna.json"

	// BINARY : executable file name
	BINARY = "_"
)

func main() {
	config, err := protocol.Load(FILENAME)

	if err != nil {
		return
	}

	argv := os.Args[1:]

	if len(argv) < 1 {
		if config == nil {
			console.Message("Not configured\n", nil)
			fmt.Println("# " + console.BOLD + fmt.Sprintf("%-12s", "init, i") + console.DEFAULT + "Create `dna.json` template")
		} else {
			console.Message("Configured!\n", nil)
		}

		fmt.Println("# " + console.BOLD + fmt.Sprintf("%-12s", "list, ls") + console.DEFAULT + "List all project scripts")
		fmt.Println("# " + console.BOLD + fmt.Sprintf("%-12s", "version, v") + console.DEFAULT + VERSION + " (" + runtime.GOOS + ")")
		fmt.Println()
		return
	}

	switch argv[0] {
	case "init", "i":
		if config == nil {
			cli.Initialize(FILENAME)
		} else {
			message := "`" + FILENAME + "` already exists"
			console.Message(message, nil)
		}

	case "list", "ls":
		if notFound(config) {
			return
		}

		count := fmt.Sprint(len(config.Scripts)) + " scripts"
		console.Message(count, nil)

		if len(config.Scripts) < 1 {
			return
		}

		fmt.Println()

		for name, script := range config.Scripts {
			fmt.Println("# " + console.BOLD + fmt.Sprintf("%-12s", name) + console.DEFAULT + script.Info)
		}

		fmt.Println()

	case "version", "v":
		version := "version " + VERSION + " (" + runtime.GOOS + ")"
		console.Message(version, nil)

	default:
		if notFound(config) {
			return
		}

		task := argv[0]
		script, exists := config.Scripts[task]

		if !exists {
			console.Error("`" + task + "` not defined")
		} else {
			err := cli.ExecSync(&task, &script)

			if err != nil {
				console.Error(err.Error())
			}
		}
	}
}

func notFound(config *protocol.DNAFile) bool {
	if config == nil {
		console.Error("`" + FILENAME + "` not found")
		return true
	}

	return false
}
