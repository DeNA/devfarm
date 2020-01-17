package awscli

import (
	"testing"
)

func TestNewVersionGetter(t *testing.T) {
	stdout := []byte("aws-cli/1.16.185 Python/3.7.4 Darwin/18.7.0 botocore/1.12.175\n")
	execute := StubExecutor(stdout, []byte{}, nil)

	getVersion := NewVersionGetter(execute)
	got, err := getVersion()

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	expected := Version(`aws-cli/1.16.185 Python/3.7.4 Darwin/18.7.0 botocore/1.12.175`)
	if got != expected {
		t.Errorf("got (%q, nil), want (%q, nil)", got, expected)
		return
	}
}
