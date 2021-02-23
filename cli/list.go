package cli

import (
	"fmt"
	"regexp"
	"sort"
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

	keys := make([]string, 0, len(*scripts))

	for key := range *scripts {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		var params []string
		script := (*scripts)[key]

		for _, template := range script.Commands {
			re := regexp.MustCompile(replace)
			params = append(params, re.FindAllString(template, -1)...)
		}

		definition := key + " " + strings.Join(params, " ")
		fmt.Println("\n+ " + console.BOLD + definition + "\n  " + console.DEFAULT + script.Info)
	}

	fmt.Println()
}
