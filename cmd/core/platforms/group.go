package platforms

import (
	"errors"
)

type InstanceGroupState string

type InstanceGroup struct {
	Name InstanceGroupName
}

func (g InstanceGroup) Less(another InstanceGroup) bool {
	return g.Name < another.Name

}

func NewInstanceGroup(name InstanceGroupName) InstanceGroup {
	return InstanceGroup{
		Name: name,
	}
}

func NewErrorInstanceGroup() InstanceGroup {
	return NewInstanceGroup("ERROR")
}

type InstanceGroupName string

func NewInstanceGroupName(s string) (InstanceGroupName, error) {
	if len(s) < 1 {
		return InstanceGroupName(""), errors.New("must not be empty")
	}

	return InstanceGroupName(s), nil
}

func NewErrorInstanceGroupName() InstanceGroupName {
	return InstanceGroupName("ERROR")
}
