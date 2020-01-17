package devicefarm

import (
	"encoding/json"
	"fmt"
)

type Run struct {
	ARN           RunARN        `json:"arn"`
	Status        RunStatus     `json:"status"`
	Result        RunResult     `json:"result"`
	DevicePoolARN DevicePoolARN `json:"devicePoolArn"`
}

func NewRun(runARN RunARN, status RunStatus, result RunResult, devicePoolARN DevicePoolARN) Run {
	return Run{
		ARN:           runARN,
		Status:        status,
		Result:        result,
		DevicePoolARN: devicePoolARN,
	}
}

type RunARN string

func (a *RunARN) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	*a = RunARN(s)
	return nil
}

func (a RunARN) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(a))
}

type RunStatus string

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
	RunStatusIsPending            RunStatus = "PENDING"
	RunStatusIsPendingConcurrency RunStatus = "PENDING_CONCURRENCY"
	RunStatusIsPendingDevice      RunStatus = "PENDING_DEVICE"
	RunStatusIsProcessing         RunStatus = "PROCESSING"
	RunStatusIsScheduling         RunStatus = "SCHEDULING"
	RunStatusIsPreparing          RunStatus = "PREPARING"
	RunStatusIsRunning            RunStatus = "RUNNING"
	RunStatusIsCompleted          RunStatus = "COMPLETED"
	RunStatusIsStopping           RunStatus = "STOPPING"
)

func (r RunStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(r))
}

func (r *RunStatus) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	switch s {
	case "PENDING":
		*r = RunStatusIsPending
	case "PENDING_CONCURRENCY":
		*r = RunStatusIsPendingConcurrency
	case "PENDING_DEVICE":
		*r = RunStatusIsPendingDevice
	case "PROCESSING":
		*r = RunStatusIsProcessing
	case "SCHEDULING":
		*r = RunStatusIsScheduling
	case "PREPARING":
		*r = RunStatusIsPreparing
	case "RUNNING":
		*r = RunStatusIsRunning
	case "COMPLETED":
		*r = RunStatusIsCompleted
	case "STOPPING":
		*r = RunStatusIsStopping
	default:
		return fmt.Errorf("unknown run status: %q", s)
	}
	return nil
}

type RunResult string

const (
	// PENDING: A pending condition.
	// PASSED: A passing condition.
	// WARNED: A warning condition.
	// FAILED: A failed condition.
	// SKIPPED: A skipped condition.
	// ERRORED: An error condition.
	// STOPPED: A stopped condition.
	RunResultIsPending RunResult = "PENDING"
	RunResultIsPassed  RunResult = "PASSED"
	RunResultIsWarned  RunResult = "WARNED"
	RunResultIsFailed  RunResult = "FAILED"
	RunResultIsSkipped RunResult = "SKIPPED"
	RunResultIsErrored RunResult = "ERRORED"
	RunResultIsStopped RunResult = "STOPPED"
)

func (r RunResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(r))
}

func (r *RunResult) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	switch s {
	case "PENDING":
		*r = RunResultIsPending
	case "PASSED":
		*r = RunResultIsPassed
	case "WARNED":
		*r = RunResultIsWarned
	case "FAILED":
		*r = RunResultIsFailed
	case "SKIPPED":
		*r = RunResultIsSkipped
	case "ERRORED":
		*r = RunResultIsErrored
	case "STOPPED":
		*r = RunResultIsStopped
	default:
		return fmt.Errorf("unknown run result: %q", s)
	}
	return nil
}
