package foreverall

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"github.com/dena/devfarm/cmd/internal/pkg/testutil"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCommand(t *testing.T) {
	stderrSpy := testutil.NewWriteCloserSpy(nil)
	procInout := cli.AnyProcInout()
	procInout.Stderr = stderrSpy

	planFile, planFileErr := pathToPlanfile()
	if planFileErr != nil {
		t.Errorf("precond failure: %s", planFileErr.Error())
		return
	}

	got := Command([]string{"--dry-run", planFile}, procInout)

	if got != cli.ExitNormal {
		t.Error(stderrSpy.Captured.String())
		return
	}
}

func pathToPlanfile() (string, error) {
	tmpDir, tmpDirErr := ioutil.TempDir(os.TempDir(), "devfarm-launch")
	if tmpDirErr != nil {
		return "", fmt.Errorf("ioutil.TempDir(_, _) == (_, %v)", tmpDirErr.Error())
	}

	path := filepath.Join(tmpDir, "planfile.yml")
	content := `instance_groups: {}`

	if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
		return "", err
	}

	return path, nil
}
