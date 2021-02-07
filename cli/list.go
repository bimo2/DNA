package cli

import (
	"fmt"

	"github.com/bimo2/DNA/console"
	"github.com/bimo2/DNA/protocol"
)

// List : print all user defined commands
func List(scripts *map[string]protocol.DNAScript) {
	count := fmt.Sprint(len(*scripts)) + " scripts"
	console.Message(count, nil)

	if len(*scripts) < 1 {
		return
	}

	fmt.Println()

	for name, script := range *scripts {
		fmt.Println("# " + console.BOLD + fmt.Sprintf("%-14s", name) + console.DEFAULT + script.Info)
	}

	fmt.Println()
}
