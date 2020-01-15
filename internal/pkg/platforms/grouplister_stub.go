package platforms

import "github.com/dena/devfarm/internal/pkg/testutil"

func AnyInstanceGroupLister() InstanceGroupLister {
	return StubInstanceGroupLister([]InstanceGroupListEntry{}, testutil.AnyError)
}

func StubInstanceGroupLister(entries []InstanceGroupListEntry, err error) InstanceGroupLister {
	return func() ([]InstanceGroupListEntry, error) {
		return entries, err
	}
}
