package runall

import (
	"github.com/dena/devfarm/cmd/core/cli"
	"github.com/dena/devfarm/cmd/core/cli/planfile"
	"github.com/dena/devfarm/cmd/core/platforms"
	"github.com/dena/devfarm/cmd/core/testutil"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	stderrSpy := testutil.NewWriteCloserSpy(nil)
	procInout := cli.AnyProcInout()
	procInout.Stderr = stderrSpy

	planfilePath, planfileErr := anyPlanfilePath()
	if planfileErr != nil {
		t.Error(planfileErr)
		return
	}

	args := []string{"--dry-run", planfilePath}
	_ = command(args, procInout)

	// FIXME: dry-run but failed now... dry-run should succeed if the given data aws correct.
	t.Log(stderrSpy.Captured.String())
}

func anyPlanfilePath() (string, error) {
	iosPlan, iosErr := anyIOSPlan()
	if iosErr != nil {
		return "", iosErr
	}

	androidPlan, androidErr := anyAndroidPlan()
	if androidErr != nil {
		return "", androidErr
	}

	plans := planfile.NewPlanfile(iosPlan, androidPlan)

	dirname, dirErr := ioutil.TempDir(os.TempDir(), "fixture")
	if dirErr != nil {
		return "", dirErr
	}

	filename := filepath.Join(dirname, "planfile.yml")
	file, openErr := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if openErr != nil {
		return "", openErr
	}
	defer file.Close()

	if writeErr := planfile.Encode(*plans, file); writeErr != nil {
		return "", writeErr
	}

	return filename, nil
}

func anyAndroidPlan() (platforms.EitherPlan, error) {
	dirname, dirErr := ioutil.TempDir(os.TempDir(), "fixture")
	if dirErr != nil {
		return platforms.EitherPlan{}, dirErr
	}

	filename := filepath.Join(dirname, "devfarm-example.apk")
	file, openErr := os.OpenFile(filename, os.O_CREATE, 0644)
	if openErr != nil {
		return platforms.EitherPlan{}, openErr
	}
	defer file.Close()

	return platforms.NewAndroidPlan(
		"any-platform",
		"any-group",
		platforms.AndroidDevice{
			DeviceName: "apple iphone xs",
			OSVersion:  "12.0",
		},
		platforms.APKPathOnLocal(filename),
		"com.example.app.id",
		platforms.AndroidIntentExtras{},
		10*time.Second,
		"test-2",
	).Either(), nil
}

func anyIOSPlan() (platforms.EitherPlan, error) {
	dirname, dirErr := ioutil.TempDir(os.TempDir(), "fixture")
	if dirErr != nil {
		return platforms.EitherPlan{}, dirErr
	}

	filename := filepath.Join(dirname, "devfarm-example.ipa")
	file, openErr := os.OpenFile(filename, os.O_CREATE, 0644)
	if openErr != nil {
		return platforms.EitherPlan{}, openErr
	}
	defer file.Close()

	return platforms.NewIOSPlan(
		"any-platform",
		"any-group",
		platforms.IOSDevice{
			DeviceName: "apple iphone xs",
			OSVersion:  "12.0",
		},
		platforms.IPAPathOnLocal(filename),
		platforms.IOSArgs{},
		10*time.Second,
		"test-1",
	).Either(), nil
}
