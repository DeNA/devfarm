package executor

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/logging"
	"io"
	"os"
)

// NOTE: You can add more interfaces if you use other File's methods.
type FileLike interface {
	io.ReadWriteCloser
	io.Seeker
	io.ReaderAt
	Stat() (os.FileInfo, error)
}

type FileOpener func(path string, flag int, perm os.FileMode) (FileLike, error)

func NewFileOpener(logger logging.SeverityLogger, dryRun bool) FileOpener {
	return func(path string, flag int, perm os.FileMode) (FileLike, error) {
		logger.Debug(fmt.Sprintf("open: %q", path))

		if flag != os.O_RDONLY && dryRun {
			logger.Debug(fmt.Sprintf("open (dry run): assume success"))
			return AnySuccessfulFile(), nil
		}

		file, err := os.OpenFile(path, flag, perm)
		if err != nil {
			logger.Debug(fmt.Sprintf("open: failed to open: %s", err.Error()))
			return nil, err
		}

		return file, nil
	}
}
