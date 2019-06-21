package devicefarm

import (
	"encoding/json"
	"math"
	"time"
)

type JobTimeout time.Duration

func (t JobTimeout) MarshalJSON() ([]byte, error) {
	return json.Marshal(math.Ceil(time.Duration(t).Minutes()))
}

type ExecutionConfiguration struct {
	JobTimeout         JobTimeout `json:"jobTimeoutMinutes"`
	AccountsCleanup    bool       `json:"accountsCleanup"`
	AppPackagesCleanup bool       `json:"appPackagesCleanup"`
	VideoCapture       bool       `json:"videoCapture"`
	SkipAppResign      bool       `json:"skipAppResign"`
}

func NewExecutionConfiguration(
	jobTimeout JobTimeout,
	accountsCleanup bool,
	appPackagesCleanup bool,
	videoCapture bool,
	skipAppResign bool,
) ExecutionConfiguration {
	return ExecutionConfiguration{
		JobTimeout:         jobTimeout,
		AccountsCleanup:    accountsCleanup,
		AppPackagesCleanup: appPackagesCleanup,
		VideoCapture:       videoCapture,
		SkipAppResign:      skipAppResign,
	}
}
