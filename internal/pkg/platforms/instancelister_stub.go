package platforms

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
)

func AnyInstanceLister() InstanceLister {
	return StubInstanceLister([]InstanceOrError{}, testutil.AnyError)
}

func StubInstanceLister(entries []InstanceOrError, err error) InstanceLister {
	return func(InstanceGroupName, InstanceListerBag) ([]InstanceOrError, error) {
		return entries, err
	}
}

func AnyInstanceListerBag() InstanceListerBag {
	return AnyBag()
}
