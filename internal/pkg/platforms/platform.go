package platforms

type Platform interface {
	ID() ID
	AuthStatusChecker() AuthStatusChecker
	DeviceLister() DeviceLister
	InstanceLister() InstanceLister
	AllInstanceLister() AllInstanceLister
	InstanceGroupLister() InstanceGroupLister
	Forever() Forever
	Runner() Runner
	IOSForever() IOSForever
	IOSRunner() IOSRunner
	AndroidForever() AndroidForever
	AndroidRunner() AndroidRunner
	Halt() Halt
	DeviceFinder() DeviceFinder
	PlanValidator() PlanValidator
}

type ID string
