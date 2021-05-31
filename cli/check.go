package cli

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/bimo2/DNA/console"
)

// Check : find system commands and files
func Check(files *map[string]string, path *string) {
	keys := make([]string, 0, len(*files))

	for key := range *files {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		var cmd *exec.Cmd
		found := true
		target := (*files)[key]

		if strings.HasPrefix(target, "/") {
			check := "test -f " + target
			cmd = exec.Command("sh", "-c", check)
		} else {
			check := "which -s " + target
			cmd = exec.Command("sh", "-c", check)
		}

		cmd.Dir = *path
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			found = false
		}

		if found {
			fmt.Println(console.BOLD + console.GREEN + "\u2713 " + key + console.DEFAULT)
		} else {
			fmt.Println(console.BOLD + console.RED + "\u2717 " + key + console.REGULAR + " - `" + target + "` not found" + console.DEFAULT)
		}
	}
}
