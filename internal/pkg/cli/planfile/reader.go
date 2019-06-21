package planfile

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"os"
)

func Read(filePath FilePath, open executor.FileOpener) (Planfile, error) {
	file, openErr := open(string(filePath), os.O_RDONLY, 0)
	if openErr != nil {
		return Planfile{}, openErr
	}
	defer file.Close()
	return Decode(file)
}
