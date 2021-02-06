package main

import (
	"fmt"
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
	dnaFile, err := cli.Load(FILENAME)

	if err != nil {
		return
	}

	argv := os.Args[1:]

	if len(argv) < 1 {
		if dnaFile == nil {
			notFound()
			return
		}

		console.Message("Configured!")
	}

	switch argv[0] {
	case "version", "v":
		console.Message("(MIT) version " + console.BOLD + VERSION)

	case "init", "i":
		if dnaFile == nil {
			cli.Initialize(FILENAME)
			return
		}

		console.Error("\"" + FILENAME + "\" file already exists")

	case "list", "ls":
		if dnaFile == nil {
			notFound()
			return
		}

		console.Message(fmt.Sprint(len(dnaFile.Scripts)) + " scripts")

		for name, script := range dnaFile.Scripts {
			fmt.Println("# " + console.BLUE + console.BOLD + fmt.Sprintf("%-12s", name) + console.RESET + " " + script.Info)
		}

	default:
		if dnaFile == nil {
			notFound()
			return
		}

		task := argv[0]
		script, exists := dnaFile.Scripts[task]

		if !exists {
			console.Error("\"" + task + "\" not defined")
		} else {
			cli.ExecSync(&script)
		}
	}
}

func notFound() {
	console.Message("Add DNA: " + console.RESET + "> " + BINARY + " init")
}
