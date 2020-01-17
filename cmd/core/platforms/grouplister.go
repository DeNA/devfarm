package platforms

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

type InstanceGroupLister func() ([]InstanceGroupListEntry, error)
