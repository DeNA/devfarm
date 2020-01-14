package exec

import (
	"bufio"
	"context"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/logging"
	"io"
	"os/exec"
)

type InteractiveExecutor func(ctx context.Context, request InteractiveRequest) error

func NewInteractiveExecutor(logger logging.SeverityLogger, dryRun bool) InteractiveExecutor {
	return func(ctx context.Context, request InteractiveRequest) error {
		if dryRun {
			return dryExecuteInteractively(logger, request)
		}
		return executeInteractively(ctx, logger, request)
	}
}

type InteractiveRequest struct {
	Stdin   io.ReadCloser
	Stdout  io.Writer
	Stderr  io.Writer
	Command string
	Args    []string
}

func NewInteractiveRequest(
	stdin io.ReadCloser,
	stdout io.Writer,
	stderr io.Writer,
	command string,
	args []string,
) InteractiveRequest {
	return InteractiveRequest{
		Stdin:   stdin,
		Stdout:  stdout,
		Stderr:  stderr,
		Command: command,
		Args:    args,
	}
}

func dryExecuteInteractively(logger logging.SeverityLogger, request InteractiveRequest) error {
	logger.Debug(fmt.Sprintf("interactive (dry run): %s", consoleLikeLine(request.Command, request.Args)))
	logger.Debug("stdin: nil\nstdout: nil\nstderr: nil\nerr: nil (assume success)")

	return nil
}

func executeInteractively(ctx context.Context, logger logging.SeverityLogger, request InteractiveRequest) error {
	logger.Debug(fmt.Sprintf("interactive: %s", consoleLikeLine(request.Command, request.Args)))

	cmd := exec.CommandContext(ctx, request.Command, request.Args...)

	stdinWriter, stdinErr := cmd.StdinPipe()
	if stdinErr != nil {
		return stdinErr
	}

	stdoutReader, stdoutErr := cmd.StdoutPipe()
	if stdoutErr != nil {
		return stdoutErr
	}
	stderrReader, stderrErr := cmd.StderrPipe()
	if stderrErr != nil {
		return stderrErr
	}

	go func() {
		scanner := bufio.NewScanner(io.TeeReader(stdoutReader, request.Stdout))
		for scanner.Scan() {
			logger.Debug(fmt.Sprintf(`stdout: %q`, scanner.Text()+"\n"))
		}
		logger.Debug("interactive: stdout closed")
	}()

	go func() {
		scanner := bufio.NewScanner(io.TeeReader(stderrReader, request.Stderr))
		for scanner.Scan() {
			logger.Debug(fmt.Sprintf(`stderr: %q`, scanner.Text()+"\n"))
		}
		logger.Debug("interactive: stderr closed")
	}()

	if err := cmd.Start(); err != nil {
		return err
	}

	defer request.Stdin.Close()
	go func() {
		defer stdinWriter.Close()

		scanner := bufio.NewScanner(io.TeeReader(request.Stdin, stdinWriter))
		for scanner.Scan() {
			logger.Debug(fmt.Sprintf(`stdin: %q`, scanner.Text()+"\n"))
		}
		logger.Debug("interactive: stdin closed")
	}()

	if err := cmd.Wait(); err != nil {
		logger.Debug(fmt.Sprintf("interactive: command %q exited with %q", request.Command, err.Error()))
		return err
	}

	return nil
}
