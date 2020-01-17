package platforms

type Instance struct {
	Device EitherDevice
	State  InstanceState
}

func NewInstance(device EitherDevice, state InstanceState) Instance {
	return Instance{
		Device: device,
		State:  state,
	}
}

func (i Instance) Less(another Instance) bool {
	if i.Device != another.Device {
		return i.Device.Less(another.Device)
	}
	return i.State.Less(another.State)
}

type InstanceState string

const (
	InstanceStateIsActivating   InstanceState = "ACTIVATING"
	InstanceStateIsActive       InstanceState = "ACTIVE"
	InstanceStateIsInactivating InstanceState = "INACTIVATING"
	InstanceStateIsInactive     InstanceState = "INACTIVE"
	InstanceStateIsUnknown      InstanceState = "UNKNOWN"
)

// FIXME: Please teach better approach to me (Kuniwak) ...
var instanceStateOrderTable = map[InstanceState]int{
	InstanceStateIsActivating:   1,
	InstanceStateIsActive:       2,
	InstanceStateIsInactivating: 3,
	InstanceStateIsInactive:     4,
	InstanceStateIsUnknown:      5,
}

func (s InstanceState) Less(another InstanceState) bool {
	return instanceStateOrderTable[s] < instanceStateOrderTable[another]
}
