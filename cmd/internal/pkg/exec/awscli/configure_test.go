package awscli

import (
	"testing"
)

func TestNewConfigStatusGetter(t *testing.T) {
	isConfigured := NewConfigStatusGetter(AnySuccessfulExecutor())

	err := isConfigured()

	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
}
