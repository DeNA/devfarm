package devicefarm

import (
	"encoding/json"
	"fmt"
)

type JobARN string

type Job struct {
	ARN    JobARN    `json:"arn"`
	Status JobStatus `json:"status"`
	Device Device    `json:"device"`
}

func NewJob(jobARN JobARN, status JobStatus, device Device) Job {
	return Job{
		ARN:    jobARN,
		Status: status,
		Device: device,
	}
}

type JobStatus string

const (
	// PENDING: A pending status.
	// PENDING_CONCURRENCY: A pending concurrency status.
	// PENDING_DEVICE: A pending device status.
	// PROCESSING: A processing status.
	// SCHEDULING: A scheduling status.
	// PREPARING: A preparing status.
	// RUNNING: A running status.
	// COMPLETED: A completed status.
	// STOPPING: A stopping status.
	JobStatusIsPending            JobStatus = "PENDING"
	JobStatusIsPendingConcurrency JobStatus = "PENDING_CONCURRENCY"
	JobStatusIsPendingDevice      JobStatus = "PENDING_DEVICE"
	JobStatusIsProcessing         JobStatus = "PROCESSING"
	JobStatusIsScheduling         JobStatus = "SCHEDULING"
	JobStatusIsPreparing          JobStatus = "PREPARING"
	JobStatusIsRunning            JobStatus = "RUNNING"
	JobStatusIsCompleted          JobStatus = "COMPLETED"
	JobStatusIsStopping           JobStatus = "STOPPING"
)

func (r JobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(r))
}

func (r *JobStatus) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	switch s {
	case "PENDING":
		*r = JobStatusIsPending
	case "PENDING_CONCURRENCY":
		*r = JobStatusIsPendingConcurrency
	case "PENDING_DEVICE":
		*r = JobStatusIsPendingDevice
	case "PROCESSING":
		*r = JobStatusIsProcessing
	case "SCHEDULING":
		*r = JobStatusIsScheduling
	case "PREPARING":
		*r = JobStatusIsPreparing
	case "RUNNING":
		*r = JobStatusIsRunning
	case "COMPLETED":
		*r = JobStatusIsCompleted
	case "STOPPING":
		*r = JobStatusIsStopping
	default:
		return fmt.Errorf("unknown run status: %q", s)
	}
	return nil
}
