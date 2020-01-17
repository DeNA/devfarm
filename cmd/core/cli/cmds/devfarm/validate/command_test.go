package validate

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/testutil"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCommand(t *testing.T) {
	stderrSpy := testutil.NewWriteCloserSpy(nil)
	procInout := cli.AnyProcInout()
	procInout.Stderr = stderrSpy

	planfile, planfileErr := pathToPlanfile()
	if planfileErr != nil {
		t.Errorf("precond failure: %s", planfileErr.Error())
		return
	}

	got := Command([]string{planfile}, procInout)

	if got != cli.ExitNormal {
		t.Errorf("got %v, want %v", got, cli.ExitNormal)
		t.Log(stderrSpy.Captured.String())
		return
	}
}

func pathToPlanfile() (string, error) {
	tmpDir, tmpDirErr := ioutil.TempDir(os.TempDir(), "devfarm-forever")
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
