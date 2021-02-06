package cli

import (
	"os"
	"os/exec"

	"github.com/bimo2/DNA/console"
)

// ExecSync : perform synchronous command
func ExecSync(script *DNAScript) error {
	for _, command := range script.Commands {
		console.Message(command)

		cmd := exec.Command("sh", "-c", command)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			return err
		}
	}

	return nil
}
