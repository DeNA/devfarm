package executor

import (
	"bytes"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/logging"
	"os/exec"
)

type Executor func(request Request) (Result, error)

func NewExecutor(logger logging.SeverityLogger, dryRun bool) Executor {
	return func(request Request) (Result, error) {
		if dryRun && !request.NoSideEffect {
			return dryExecute(logger, request.Command, request.Args)
		}

		return execute(logger, request.Command, request.Args)
	}
}

type Request struct {
	Command      string
	Args         []string
	NoSideEffect bool
}

func newRequest(command string, args []string, noSideEffect bool) Request {
	return Request{
		Command:      command,
		Args:         args,
		NoSideEffect: noSideEffect,
	}
}

func NewRequest(command string, args []string) Request {
	return newRequest(command, args, false)
}

type Result struct {
	Stdout []byte
	Stderr []byte
}

func NewResult(stdout []byte, stderr []byte) Result {
	return Result{
		Stdout: stdout,
		Stderr: stderr,
	}
}

func dryExecute(logger logging.SeverityLogger, command string, args []string) (Result, error) {
	logger.Debug(fmt.Sprintf("exec (dry run): %s", consoleLikeLine(command, args)))

	logger.Debug("stdin: (empty)")
	logger.Debug("stdout:\nnil\nstderr:nil\nerr: nil (assume success)")

	return emptyResult(), nil
}

func execute(logger logging.SeverityLogger, command string, args []string) (Result, error) {
	logger.Debug(fmt.Sprintf("exec: %s", consoleLikeLine(command, args)))

	cmd := exec.Command(command, args...)

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		return emptyResult(), err
	}

	if err := cmd.Wait(); err != nil {
		logExecutionResult(logger, stdout.String(), stderr.String(), err)
		return emptyResult(), fmt.Errorf("exec: command %q exited with: %q\n%s", command, err.Error(), stderrHint(stderr.String()))
	}

	logExecutionResult(logger, stdout.String(), stderr.String(), nil)
	return NewResult(stdout.Bytes(), stderr.Bytes()), nil
}

func emptyResult() Result {
	return NewResult([]byte{}, []byte{})
}
