package exec

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func FakeFileOpener(content []byte) FileOpener {
	return func(path string, flag int, perm os.FileMode) (FileLike, error) {
		dirname, dirErr := ioutil.TempDir(os.TempDir(), "devfarm-fake-file")
		if dirErr != nil {
			return nil, dirErr
		}

		fakePath := filepath.Join(dirname, filepath.Base(path))

		fakeFile, openErr := os.OpenFile(fakePath, os.O_WRONLY|os.O_CREATE, 0644)
		if openErr != nil {
			return nil, openErr
		}

		if _, err := fakeFile.Write(content); err != nil {
			return nil, err
		}

		if err := fakeFile.Close(); err != nil {
			return nil, err
		}

		return os.OpenFile(fakePath, flag, perm)
	}
}
