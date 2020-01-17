package exec

import (
	"errors"
	"fmt"
	"github.com/dena/devfarm/cmd/core/logging"
	"os"
	"os/exec"
	"strings"
)

type ExecutableFinder func(string) error

func NewExecutableFinder(logger logging.SeverityLogger, dryRun bool) ExecutableFinder {
	if dryRun {
		return func(name string) error {
			logger.Debug(fmt.Sprintf("find (dry run): %q", name))
			logger.Debug(fmt.Sprintf("find (assume success): /path/to/%q", name))
			return nil
		}
	}

	return func(name string) error {
		return FindExecutable(logger, name)
	}
}

func FindExecutable(logger logging.SeverityLogger, name string) error {
	logger.Debug(fmt.Sprintf("find: %q", name))

	path, err := exec.LookPath(name)
	if err != nil {
		pathEnv := os.Getenv("PATH")

		var note string
		if strings.Contains(pathEnv, "~") {
			note = "\n\n" + NoteAboutTilde
		}

		message := fmt.Sprintf(`PATH=%s%s`, pathEnv, note)
		logger.Debug(fmt.Sprintf("find (failed): %s", message))
		return errors.New(message)
	}

	logger.Debug(fmt.Sprintf("find (success): %q", path))
	return nil
}

var NoteAboutTilde = `NOTE: "~/" is not recognized as "$HOME/". So if you using "~/" in your PATH, you should replace "~" with "$HOME".`
