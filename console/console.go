package console

import (
	"fmt"
)

// Message : print program messages
func Message(s string) {
	fmt.Println("\u001b[34m[DNA] " + s + "\u001b[0m")
}

// Error : print program errors
func Error(s string) {
	fmt.Println("\u001b[31m[DNA] " + s + "\u001b[0m")
}
