package main

import (
	"fmt"
	"os"
)

const (
	// VERSION : DNA version
	VERSION = "0.1.0"

	// BINARY : executable file name
	BINARY = "_"
)

func main() {
	argv := os.Args[1:]

	fmt.Println(argv)

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
