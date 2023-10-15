package ashpalm

import (
	"bytes"
	"log/slog"
	"os/exec"

	mymazda "github.com/taylormonacelli/forestfish/mymazda"
)

func RunCmd(cmd *exec.Cmd, cwd string) (int, string, string) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = cwd

	if !mymazda.DirExists(cwd) {
		slog.Error("can't change to directory", "directory", cwd)
		return 1, "", ""
	}

	err := cmd.Run()

	if err == nil {
		slog.Debug("command executed successfully", "cmd", cmd.String())
		return 0, stdout.String(), stderr.String()
	}

	exitErr, ok := err.(*exec.ExitError)
	if ok {
		slog.Error("command had error", "cmd", cmd.String(), "code", exitErr.ExitCode())
		return exitErr.ExitCode(), stdout.String(), stderr.String()
	}

	slog.Error("not ok", "error", err.Error())

	return 1, stdout.String(), stderr.String()
}
