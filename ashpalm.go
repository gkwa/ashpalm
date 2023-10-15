package ashpalm

import (
	"bytes"
	"log/slog"
	"os/exec"
)

func RunCmd(cmd *exec.Cmd) int {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err == nil {
		slog.Debug("command executed successfully", "cmd", cmd.String())
		return 0
	}

	slog.Debug("stdout", "cmd", cmd.String(), "str", stdout.String())
	slog.Debug("stderr", "cmd", cmd.String(), "str", stderr.String())

	exitErr, ok := err.(*exec.ExitError)
	if ok {
		slog.Error("command had error", "cmd", cmd.String(), "code", exitErr.ExitCode())
		return exitErr.ExitCode()
	}

	slog.Error("not ok", "error", err.Error())

	return 1
}
