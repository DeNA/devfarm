package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/internal/pkg/platforms"
)

func instanceStateFrom(status devicefarm.JobStatus) (platforms.InstanceState, error) {
	switch status {
	case devicefarm.JobStatusIsPending, devicefarm.JobStatusIsPendingConcurrency, devicefarm.JobStatusIsPendingDevice,
		devicefarm.JobStatusIsProcessing, devicefarm.JobStatusIsScheduling, devicefarm.JobStatusIsPreparing:
		return platforms.InstanceStateIsActivating, nil

	case devicefarm.JobStatusIsRunning:
		return platforms.InstanceStateIsActive, nil

	case devicefarm.JobStatusIsCompleted:
		return platforms.InstanceStateIsInactive, nil

	case devicefarm.JobStatusIsStopping:
		return platforms.InstanceStateIsInactivating, nil

	default:
		return platforms.InstanceStateIsUnknown, fmt.Errorf("unknown run state: %v", status)
	}
}
