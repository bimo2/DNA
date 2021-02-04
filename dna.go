package main

import (
	"fmt"
	"os"

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

	fmt.Println(argv, dnaFile)

	if len(argv) < 1 {
		return
	}

	switch argv[0] {
	case "--version", "-v":
		fmt.Println("DNA version " + VERSION)

	default:
		return
	}
}
