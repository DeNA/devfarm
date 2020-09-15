package awscli

import (
	"github.com/dena/devfarm/cmd/core/exec"
	"testing"
)

func TestNewConfigStatusGetterViaEnv(t *testing.T) {
	isConfigured := NewConfigStatusGetter(AnyFailedExecutor(), func(envName string) string {
		switch envName {
		case "AWS_ACCESS_KEY_ID":
			return "********************"
		case "AWS_SECRET_ACCESS_KEY":
			return "*************/*******/******************"
		default:
			panic(envName)
		}
	})

	err := isConfigured()

	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
}

func TestNewConfigStatusGetterViaAwsCli(t *testing.T) {
	isConfigured := NewConfigStatusGetter(AnySuccessfulExecutor(), exec.StubEnvGetter(""))

	err := isConfigured()

	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
}
