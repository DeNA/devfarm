package exec

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
	"os"
)

func AnySuccessfulFileOpener() FileOpener {
	return StubFileOpener(AnySuccessfulFile(), nil)
}

func AnyFailedFileOpener() FileOpener {
	return StubFileOpener(nil, testutil.AnyError)
}

func StubFileOpener(file FileLike, err error) FileOpener {
	return func(string, int, os.FileMode) (FileLike, error) {
		return file, err
	}
}
