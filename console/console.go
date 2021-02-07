package console

import (
	"fmt"
)

const (
	// DEFAULT : default console font interpolation
	DEFAULT = "\u001b[0m"

	// BOLD : console font style interpolation
	BOLD = "\u001b[1m"

	// REGULAR : console font style interpolation
	REGULAR = "\u001b[22m"

	// RED : console font colour interpolation
	RED = "\u001b[31m"

	// GREEN : console font colour interpolation
	GREEN = "\u001b[32m"

	// BLUE : console font colour interpolation
	BLUE = "\u001b[34m"
)

// Message : print message
func Message(log string, context *string) {
	tag := "DNA"

	if context != nil {
		tag = *context
	}

	fmt.Println(BLUE + BOLD + tag + " " + REGULAR + log + DEFAULT)
}

// Error : print error message
func Error(log string) {
	fmt.Println(RED + BOLD + "ERROR " + REGULAR + log + DEFAULT)
}

// Success : print success message
func Success(log string) {
	fmt.Println(GREEN + BOLD + "DONE " + REGULAR + log + DEFAULT)
}
