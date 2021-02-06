package cli

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/bimo2/DNA/console"
)

// ExecSync : perform synchronous command
func ExecSync(script *DNAScript) error {
	start := time.Now()

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

	elapsed := time.Now().Sub(start)
	fmt.Println("Time: " + elapsed.String())
	return nil
}
