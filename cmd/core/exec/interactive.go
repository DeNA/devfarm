package exec

import (
	"bufio"
	"context"
	"fmt"
	"github.com/dena/devfarm/cmd/core/logging"
	"io"
	"os/exec"
)

type InteractiveExecutor interface {
	Execute(ctx context.Context, request InteractiveRequest) error
}

func NewInteractiveExecutor(logger logging.SeverityLogger, dryRun bool) InteractiveExecutor {
	if dryRun {
		return &dryRunInteractiveExecutor{logger: logger}
	}
	return &interactiveExecutor{logger: logger}
}

type InteractiveRequest struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Command string
	Args    []string
}

func NewInteractiveRequest(
	stdin io.Reader,
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

type interactiveExecutor struct {
	logger logging.SeverityLogger
}

var _ InteractiveExecutor = &interactiveExecutor{}

func (i *interactiveExecutor) Execute(ctx context.Context, request InteractiveRequest) error {
	i.logger.Debug(fmt.Sprintf("interactive: %s", consoleLikeLine(request.Command, request.Args)))

	cmd := exec.CommandContext(ctx, request.Command, request.Args...)

	stdinWriter, stdinErr := cmd.StdinPipe()
	if stdinErr != nil {
		return stdinErr
	}
	i.teeWriter("stdin", stdinWriter, request.Stdin)

	stdoutReader, stdoutErr := cmd.StdoutPipe()
	if stdoutErr != nil {
		return stdoutErr
	}
	i.teeReader("stdout", stdoutReader, request.Stdout)

	stderrReader, stderrErr := cmd.StderrPipe()
	if stderrErr != nil {
		return stderrErr
	}
	i.teeReader("stderr", stderrReader, request.Stderr)

	if err := cmd.Run(); err != nil {
		i.logger.Debug(fmt.Sprintf("interactive: command %q exited with %q", request.Command, err.Error()))
		return err
	}

	return nil
}

// To print what data are read.
func (i *interactiveExecutor) teeReader(name string, reader io.Reader, writer io.Writer) {
	var teeReader io.Reader
	if writer == nil {
		teeReader = reader
	} else {
		teeReader = io.TeeReader(reader, writer)
	}

	go func() {
		scanner := bufio.NewScanner(teeReader)
		for scanner.Scan() {
			i.logger.Debug(fmt.Sprintf(`%s: %q`, name, scanner.Text()+"\n"))
		}
		i.logger.Debug(fmt.Sprintf("interactive: %s closed", name))
	}()
}

// To print what data are write.
func (i *interactiveExecutor) teeWriter(name string, writer io.WriteCloser, reader io.Reader) {
	if reader == nil {
		if err := writer.Close(); err != nil {
			i.logger.Error(fmt.Sprintf("interactive: failed to close stdin: %q", err.Error()))
		}
		return
	}

	go func() {
		defer writer.Close()

		scanner := bufio.NewScanner(io.TeeReader(reader, writer))
		for scanner.Scan() {
			i.logger.Debug(fmt.Sprintf(`%s: %q`, name, scanner.Text()+"\n"))
		}
		i.logger.Debug(fmt.Sprintf("interactive: %s closed", name))
	}()
}

type dryRunInteractiveExecutor struct {
	logger logging.SeverityLogger
}

func (d *dryRunInteractiveExecutor) Execute(_ context.Context, request InteractiveRequest) error {
	d.logger.Debug(fmt.Sprintf("interactive (dry run): %s", consoleLikeLine(request.Command, request.Args)))
	d.logger.Debug("stdin: nil\nstdout: nil\nstderr: nil\nerr: nil (assume success)")
	return nil
}

var _ InteractiveExecutor = &dryRunInteractiveExecutor{}
