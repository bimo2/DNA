package cli

import (
	"os"
	"os/exec"
	"time"

	"github.com/bimo2/DNA/console"
	"github.com/bimo2/DNA/protocol"
)

// ExecSync : perform synchronous command
func ExecSync(context *string, script *protocol.DNAScript) error {
	start := time.Now()

	for _, command := range script.Commands {
		console.Message(command, context)

		cmd := exec.Command("sh", "-c", command)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			return err
		}
	}

	elapsed := time.Now().Sub(start)
	console.Success(elapsed.String())
	return nil
}
