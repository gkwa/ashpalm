package main

import (
	"flag"
	"log/slog"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/taylormonacelli/ashpalm"
	"github.com/taylormonacelli/goldbug"
)

var (
	verbose   bool
	logFormat string
)

func main() {
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&verbose, "v", false, "Enable verbose output (shorthand)")

	flag.StringVar(&logFormat, "log-format", "", "Log format (text or json)")

	flag.Parse()

	if verbose || logFormat != "" {
		if logFormat == "json" {
			goldbug.SetDefaultLoggerJson(slog.LevelDebug)
		} else {
			goldbug.SetDefaultLoggerText(slog.LevelDebug)
		}
	}

	nonExistantDir := getNonExistantPath()

	checkDirExists(nonExistantDir)

	checkDirExists(".")
}

func getNonExistantPath() string {
	currentTime := time.Now()
	epochTime := currentTime.Unix()

	strValue := strconv.FormatInt(epochTime, 10)

	return filepath.Join("/tmp", "ashpalm"+strValue)
}

func checkDirExists(dir string) {
	cmd := exec.Command("ls", dir)
	intResult, _, _ := ashpalm.RunCmd(cmd)

	if intResult == 0 {
		slog.Debug("directory exists", "code", intResult)
	}
}
