package platforms

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func AnyAllInstanceLister() AllInstanceLister {
	return StubAllInstanceLister([]InstanceOrError{}, testutil.AnyError)
}

func StubAllInstanceLister(entries []InstanceOrError, err error) AllInstanceLister {
	return func() ([]InstanceOrError, error) {
		return entries, err
	}
}
