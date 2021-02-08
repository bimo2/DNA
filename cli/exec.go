package cli

import (
	"os"
	"os/exec"
	"time"

	"github.com/bimo2/DNA/console"
	"github.com/bimo2/DNA/protocol"
)

// ExecSync : perform synchronous command
func ExecSync(context *string, script *protocol.DNAScript, path *string) {
	start := time.Now()

	for _, command := range script.Commands {
		console.Message(command, context)

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
