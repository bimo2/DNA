package cli

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/bimo2/DNA/console"
	"github.com/bimo2/DNA/protocol"
)

const replace = "\\[\\w+\\]"

// ExecSync : perform synchronous command
func ExecSync(argv *[]string, script *protocol.DNAScript, env *map[string]string, path *string) {
	context := (*argv)[0]
	insert := 0

	useEnv := func(template string) string {
		for key, value := range *env {
			match := "&" + key
			template = strings.ReplaceAll(template, match, value)
		}

		return template
	}

	next := func(string) string {
		if insert++; insert < len(*argv) {
			return (*argv)[insert]
		}

		return ""
	}

	start := time.Now()

	for _, template := range script.Commands {
		template = useEnv(template)
		re := regexp.MustCompile(replace)
		command := re.ReplaceAllStringFunc(template, next)

		if command[:2] == "# " {
			console.Message(command[2:], &context)
			continue
		}

		console.Message(command, &context)
		cmd := exec.Command("sh", "-c", command)
		cmd.Dir = *path
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			console.Error(err.Error())
			return
		}
	}

	elapsed := time.Now().Sub(start)
	console.Success(elapsed.String())
}
