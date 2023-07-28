package howlong

import (
	"os"
	"os/exec"
	"time"
)

func Run(program string, args ...string) (time.Duration, error) {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	start := time.Now()
	if err := cmd.Run(); err != nil {
		return 0, err
	}
	return time.Since(start), nil
}
