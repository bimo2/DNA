package console

import (
	"fmt"
)

const (
	// BOLD : console bold text
	BOLD = "\033[1m"

	// BLUE : console blue text
	BLUE = "\u001b[34m"

	// RED : console red text
	RED = "\u001b[31m"

	// RESET : console default text
	RESET = "\u001b[0m"
)

// Message : print program messages
func Message(log string) {
	fmt.Println(BLUE + "[DNA] " + log + RESET)
}

// Error : print program errors
func Error(log string) {
	fmt.Println(RED + "[DNA] " + log + RESET)
}
