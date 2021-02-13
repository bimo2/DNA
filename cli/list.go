package cli

import (
	"fmt"
	"regexp"
	"strings"

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

	for name, script := range *scripts {
		var params []string

		for _, template := range script.Commands {
			re := regexp.MustCompile(replace)
			params = append(params, re.FindAllString(template, -1)...)
		}

		definition := name + " " + strings.Join(params, " ")
		fmt.Println("\n+ " + console.BOLD + definition + "\n  " + console.DEFAULT + script.Info)
	}

	fmt.Println()
}
