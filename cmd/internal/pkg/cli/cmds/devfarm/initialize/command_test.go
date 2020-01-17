package initialize

import (
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"github.com/dena/devfarm/cmd/internal/pkg/testutil"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCommand(t *testing.T) {
	stderrSpy := testutil.NewWriteCloserSpy(nil)
	procInout := cli.AnyProcInout()
	procInout.Stderr = stderrSpy

	workspace, tempDirErr := ioutil.TempDir(os.TempDir(), "devfarm-init")
	if tempDirErr != nil {
		t.Errorf("precond failure: %s", tempDirErr)
		return
	}

	filePath := path.Join(workspace, "planfile.yml")
	got := command([]string{filePath}, procInout)

	if got != cli.ExitNormal {
		t.Error(stderrSpy.Captured.String())
		return
	}
}
