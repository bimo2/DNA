package console

import (
	"fmt"
)

// Message : print program messages
func Message(log string) {
	fmt.Println("\u001b[34m[DNA] " + log + "\u001b[0m")
}

// Error : print program errors
func Error(log string) {
	fmt.Println("\u001b[31m[DNA] " + log + "\u001b[0m")
}
