package awscli

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"testing"
)

func TestNewInstallStatusGetter(t *testing.T) {
	isInstalled := NewInstallStatusGetter(executor.AnySuccessfulExecutableFinder)

	err := isInstalled()

	if err != nil {
		t.Errorf("got %v, want nil", err)
		return
	}
}
