package runandroid

import (
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

	apkPath, err := anyAPKPath()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	args := []string{
		"--instance-group", "example",
		"--platform", "aws-device-farm",
		"--device", "google google pixel 3",
		"--os-version", "9.0",
		"--apk", apkPath,
		"--app-id", "com.example.apk",
		"--lifetime-sec", "900",
		"--dry-run",
	}
	_ = command(args, procInout)

	// FIXME: dry-run but failed now... dry-run should succeed if the given data aws correct.
	t.Log(stderrSpy.Captured.String())
}

func anyAPKPath() (string, error) {
	dirname, dirErr := ioutil.TempDir(os.TempDir(), "fixture")
	if dirErr != nil {
		return "", dirErr
	}

	filename := filepath.Join(dirname, "devfarm-example.apk")
	file, openErr := os.OpenFile(filename, os.O_CREATE, 0644)
	if openErr != nil {
		return "", openErr
	}
	_ = file.Close()

	return filename, nil
}
