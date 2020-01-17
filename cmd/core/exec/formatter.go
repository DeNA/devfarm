package exec

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/logging"
	"strings"
)

func consoleLikeLine(name string, args []string) string {
	cmdAndArgs := make([]string, len(args)+1)
	cmdAndArgs[0] = name
	for i, arg := range args {
		cmdAndArgs[i+1] = arg
	}
	return fmt.Sprintf("%q", cmdAndArgs)
}

func logExecutionResult(logger logging.SeverityLogger, stdout string, stderr string, err error) {
	logger.Debug(stdoutHint(stdout))
	logger.Debug(stderrHint(stderr))
	logger.Debug(fmt.Sprintf("error: %v", err))
}

func stdoutHint(stdout string) string {
	if len(stdout) > 0 {
		return fmt.Sprintf("stdout:\n%s\n", strings.TrimSpace(stdout))
	} else {
		return "stdout: (empty)"
	}
}

func stderrHint(stderr string) string {
	if len(stderr) > 0 {
		return fmt.Sprintf("stderr:\n%s\n", strings.TrimSpace(stderr))
	} else {
		return "stderr: (empty)"
	}
}
