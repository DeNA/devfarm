package cli

import (
	"github.com/dena/devfarm/internal/pkg/logging"
	"io"
)

func NewLogger(verbose bool, out io.Writer) logging.SeverityLogger {
	if verbose {
		return logging.NewLogger(logging.Debug, out)
	}
	return logging.NewLogger(logging.Info, out)
}
