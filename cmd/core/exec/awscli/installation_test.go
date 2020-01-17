package awscli

import (
	"github.com/dena/devfarm/cmd/core/exec"
	"testing"
)

func TestNewInstallStatusGetter(t *testing.T) {
	isInstalled := NewInstallStatusGetter(exec.AnySuccessfulExecutableFinder)

	err := isInstalled()

	if err != nil {
		t.Errorf("got %v, want nil", err)
		return
	}
}
