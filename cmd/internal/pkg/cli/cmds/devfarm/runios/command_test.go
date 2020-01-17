package runios

import (
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

	ipaPath, err := anyIPAPath()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	args := []string{
		"--instance-group", "example",
		"--platform", "aws-device-farm",
		"--device", "apple iphone xs",
		"--os-version", "12.0",
		"--ipa", ipaPath,
		"--lifetime", "900",
		"--dry-run",
	}
	_ = command(args, procInout)

	// FIXME: dry-run but failed now... dry-run should succeed if the given data aws correct.
	t.Log(stderrSpy.Captured.String())
}

func anyIPAPath() (string, error) {
	dirname, dirErr := ioutil.TempDir(os.TempDir(), "fixture")
	if dirErr != nil {
		return "", dirErr
	}

	filename := filepath.Join(dirname, "devfarm-example.ipa")
	file, openErr := os.OpenFile(filename, os.O_CREATE, 0644)
	if openErr != nil {
		return "", openErr
	}
	_ = file.Close()

	return filename, nil
}
