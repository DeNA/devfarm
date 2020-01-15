package planfile

import (
	"github.com/dena/devfarm/internal/pkg/exec"
	"os"
)

func Write(planfile Planfile, filePath FilePath, open exec.FileOpener) error {
	file, openErr := open(string(filePath), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if openErr != nil {
		return openErr
	}
	defer file.Close()
	return Encode(planfile, file)
}
