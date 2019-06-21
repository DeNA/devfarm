package awsdevicefarm

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewTestPackageUploader(t *testing.T) {
	spyLogger := logging.SpySeverityLogger()
	reserveAndUploaderCallArgs, spyReserveAndUploader := spyReserveUploaderIfNotExists(anySuccessfulReserveAndUploaderIfNotExists())

	uploadTestPackage := newTestPackageUploader(spyLogger, newTestPackageGen(), platforms.NewCRC32Hasher(), spyReserveAndUploader)

	_, err := uploadTestPackage("arn:aws:devicefarm:ANY_PROJECT")
	if err != nil {
		t.Errorf("got %v, want nil", err)
		t.Log(spyLogger.Logs.String())
		return
	}

	if len(*reserveAndUploaderCallArgs) != 1 {
		t.Errorf("number of reserveAndUpload calls are %v, want 1", len(*reserveAndUploaderCallArgs))
		t.Log(spyLogger.Logs.String())
		return
	}

	size := (*reserveAndUploaderCallArgs)[0].size
	if size < 1 {
		t.Errorf("got %d, want > 0", size)
		t.Log(spyLogger.Logs.String())
		return
	}
}

func TestNewTestPackageBundleGen(t *testing.T) {
	path, filenameErr := tempTestPackageFilepath("noop-tests-0.0.0.tgz")
	if filenameErr != nil {
		t.Errorf("want nil, got %v", filenameErr)
		return
	}

	file, openErr := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if openErr != nil {
		t.Errorf("want nil, got %v", openErr)
		t.Log(path)
		return
	}
	defer file.Close()

	if err := writeTestPackageBundle(func(header *zip.FileHeader) (io.Writer, error) { return file, nil }); err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}

	if err := checkValidTgz(path); err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}
}

func TestNewTestPackageArchiveGenValidZip(t *testing.T) {
	path, filenameErr := tempTestPackageFilepath("devfarm-test-package.zip")
	if filenameErr != nil {
		t.Errorf("want nil, got %v", filenameErr)
		return
	}

	file, openErr := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if openErr != nil {
		t.Errorf("want nil, got %v", openErr)
		return
	}
	defer file.Close()

	embedded := anyTestPackageEmbeddedData()
	testPackageGen := newTestPackageGen()

	if err := testPackageGen(embedded, file); err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}

	_ = file.Sync()
	if err := checkValidZip(path); err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}
}

func TestNewTestPackageArchiveGenPurity(t *testing.T) {
	first := &bytes.Buffer{}
	embedded := anyTestPackageEmbeddedData()
	testPackageGen := newTestPackageGen()

	if err := testPackageGen(embedded, first); err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}

	for i := 0; i < 10; i++ {
		buf := &bytes.Buffer{}
		if err := testPackageGen(embedded, buf); err != nil {
			t.Errorf("want nil, got %v", err)
			return
		}

		if !reflect.DeepEqual(first.Bytes(), buf.Bytes()) {
			t.Error(cmp.Diff(first.Bytes(), buf.Bytes()))
			return
		}
	}
}

func tempTestPackageFilepath(basename string) (string, error) {
	dir, dirErr := ioutil.TempDir(os.TempDir(), "devfarm-test-package")
	if dirErr != nil {
		return "", dirErr
	}

	return filepath.Join(dir, basename), nil
}

func checkValidZip(filename string) error {
	cmd := exec.Command("zipinfo", filename)

	stderr, stderrErr := cmd.StderrPipe()
	if stderrErr != nil {
		return stderrErr
	}

	if err := cmd.Run(); err != nil {
		bytes, _ := ioutil.ReadAll(stderr)
		return fmt.Errorf("$ zipinfo %s\n%s", filename, string(bytes))
	}

	return nil
}

func checkValidTgz(filename string) error {
	cmd := exec.Command("tar", "-tvf", filename)

	stdout, stdoutErr := cmd.StdoutPipe()
	if stdoutErr != nil {
		return stdoutErr
	}

	stderr, stderrErr := cmd.StderrPipe()
	if stderrErr != nil {
		return stderrErr
	}

	if err := cmd.Run(); err != nil {
		stdoutBytes, _ := ioutil.ReadAll(stdout)
		stderrBytes, _ := ioutil.ReadAll(stderr)
		return fmt.Errorf("failed to execute `tar -tvf %s`: %s\n%q\n%q",
			filename, err.Error(), string(stdoutBytes), string(stderrBytes))
	}

	return nil
}
