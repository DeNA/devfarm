package platforms

import (
	"github.com/dena/devfarm/cmd/core/testutil"
)

func AnyInstanceLister() InstanceLister {
	return StubInstanceLister([]InstanceOrError{}, testutil.AnyError)
}

func StubInstanceLister(entries []InstanceOrError, err error) InstanceLister {
	return func(InstanceGroupName) ([]InstanceOrError, error) {
		return entries, err
	}
}
