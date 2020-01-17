package platforms

type PlatformStub struct {
	id                      ID
	AuthStatusCheckerFunc   AuthStatusChecker
	DeviceListerFunc        DeviceLister
	InstanceListerFunc      InstanceLister
	AllInstanceListerFunc   AllInstanceLister
	InstanceGroupListerFunc InstanceGroupLister
	ForeverFunc             Forever
	RunnerFunc              Runner
	IOSForeverFunc          IOSForever
	IOSRunnerFunc           IOSRunner
	AndroidForeverFunc      AndroidForever
	AndroidRunnerFunc       AndroidRunner
	HaltFunc                Halt
	DeviceFinderFunc        DeviceFinder
	PlanValidatorFunc       PlanValidator
}

var _ Platform = PlatformStub{}

func AnyPlatform() *PlatformStub {
	return &PlatformStub{
		id:                      "any",
		AuthStatusCheckerFunc:   AnyAuthStatusChecker(),
		DeviceListerFunc:        AnyDeviceLister(),
		InstanceListerFunc:      AnyInstanceLister(),
		AllInstanceListerFunc:   AnyAllInstanceLister(),
		InstanceGroupListerFunc: AnyInstanceGroupLister(),
		ForeverFunc:             AnyForever(),
		RunnerFunc:              AnyRunner(),
		IOSForeverFunc:          AnyIOSForever(),
		IOSRunnerFunc:           AnyIOSRunner(),
		AndroidForeverFunc:      AnyAndroidForever(),
		AndroidRunnerFunc:       FailedAndroidRunner(),
		HaltFunc:                AnyHalt(),
		DeviceFinderFunc:        AnyDeviceFinder(),
		PlanValidatorFunc:       AnyPlanValidator(),
	}
}

func (p PlatformStub) ID() ID {
	return p.id
}

func (p PlatformStub) AuthStatusChecker() AuthStatusChecker {
	return p.AuthStatusCheckerFunc
}

func (p PlatformStub) DeviceLister() DeviceLister {
	return p.DeviceListerFunc
}

func (p PlatformStub) InstanceLister() InstanceLister {
	return p.InstanceListerFunc
}

func (p PlatformStub) AllInstanceLister() AllInstanceLister {
	return p.AllInstanceListerFunc
}

func (p PlatformStub) InstanceGroupLister() InstanceGroupLister {
	return p.InstanceGroupListerFunc
}

func (p PlatformStub) Forever() Forever {
	return p.ForeverFunc
}

func (p PlatformStub) Runner() Runner {
	return p.RunnerFunc
}

func (p PlatformStub) IOSForever() IOSForever {
	return p.IOSForeverFunc
}

func (p PlatformStub) IOSRunner() IOSRunner {
	return p.IOSRunnerFunc
}

func (p PlatformStub) AndroidForever() AndroidForever {
	return p.AndroidForeverFunc
}

func (p PlatformStub) AndroidRunner() AndroidRunner {
	return p.AndroidRunnerFunc
}

func (p PlatformStub) Halt() Halt {
	return p.HaltFunc
}

func (p PlatformStub) DeviceFinder() DeviceFinder {
	return p.DeviceFinderFunc
}

func (p PlatformStub) PlanValidator() PlanValidator {
	return p.PlanValidatorFunc
}
