package planfile

import (
	"github.com/dena/devfarm/internal/pkg/exec"
	"os"
)

func Read(filePath FilePath, open exec.FileOpener) (Planfile, error) {
	file, openErr := open(string(filePath), os.O_RDONLY, 0)
	if openErr != nil {
		return Planfile{}, openErr
	}
	defer file.Close()
	return Decode(file)
}
