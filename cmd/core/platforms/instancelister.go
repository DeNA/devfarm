package platforms

type InstanceLister func(groupName InstanceGroupName) ([]InstanceOrError, error)

type InstanceOrError struct {
	Instance
	Error error
}

func NewInstanceListEntry(instance Instance, err error) InstanceOrError {
	return InstanceOrError{
		Instance: instance,
		Error:    err,
	}
}

func (i InstanceOrError) Less(another InstanceOrError) bool {
	if i.Instance != another.Instance {
		return i.Instance.Less(another.Instance)
	}

	if i.Error == nil {
		return false
	}

	if another.Error == nil {
		return true
	}

	return i.Error.Error() < another.Error.Error()
}
