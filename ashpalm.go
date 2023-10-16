package ashpalm

import (
	"bytes"
	"os/exec"
)

func RunCmd(cmd *exec.Cmd) (int, string, string) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err == nil {
		return 0, stdout.String(), stderr.String()
	}

	exitErr, ok := err.(*exec.ExitError)
	if ok {
		return exitErr.ExitCode(), stdout.String(), stderr.String()
	}

	return 1, stdout.String(), stderr.String()
}
