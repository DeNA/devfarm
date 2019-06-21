package platforms

import (
	"github.com/dena/devfarm/internal/pkg/executor"
	"github.com/dena/devfarm/internal/pkg/logging"
)

type InstanceGroupListerBag interface {
	GetLogger() logging.SeverityLogger
	GetExecutor() executor.Executor
	GetFinder() executor.ExecutableFinder
}

type InstanceGroupListEntry struct {
	Group      InstanceGroup
	GroupError error
}

func (i InstanceGroupListEntry) Less(another InstanceGroupListEntry) bool {
	if i.Group != another.Group {
		return i.Group.Less(another.Group)
	}

	if i.GroupError == nil {
		return true
	}

	if another.GroupError == nil {
		return false
	}

	return i.GroupError.Error() < another.GroupError.Error()
}

func NewInstanceGroupListEntry(group InstanceGroup, err error) InstanceGroupListEntry {
	return InstanceGroupListEntry{
		Group:      group,
		GroupError: err,
	}
}

type InstanceGroupLister func(bag InstanceGroupListerBag) ([]InstanceGroupListEntry, error)
