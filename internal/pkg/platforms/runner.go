package platforms

import (
	"fmt"
	"strings"
	"sync"
)

type Runner func(plans []EitherPlan) (Results, error)

func NewUnoptimizedRunner(runIOS IOSRunner, runAndroid AndroidRunner) Runner {
	return func(plans []EitherPlan) (Results, error) {
		var wg sync.WaitGroup
		results := NewResults()

		for _, plan := range plans {
			switch plan.OSName {
			case OSIsIOS:
				wg.Add(1)
				go func(iosPlan IOSPlan) {
					if iosErr := runIOS(iosPlan); iosErr != nil {
						results.AddErrorOrNils(iosErr)
					} else {
						results.AddSuccesses(1)
					}
					wg.Done()
				}(plan.IOS())

			case OSIsAndroid:
				wg.Add(1)
				go func(androidPlan AndroidPlan) {
					if androidErr := runAndroid(androidPlan); androidErr != nil {
						results.AddErrorOrNils(androidErr)
					} else {
						results.AddSuccesses(1)
					}
					wg.Done()
				}(plan.Android())

			default:
				results.AddErrorOrNils(fmt.Errorf("unknown OS: %q", plan.OSName))
			}
		}

		wg.Wait()
		return *results, results.Err()
	}
}

type RunError struct {
	Passed int
	Errors []error
}

var _ error = RunError{}

func (e RunError) Error() string {
	errCount := len(e.Errors)
	totalCount := e.Passed + errCount
	strArr := make([]string, len(e.Errors))
	for i, err := range e.Errors {
		strArr[i] = err.Error()
	}
	return fmt.Sprintf("Failed %d/%d\n:%s", errCount, totalCount, strings.Join(strArr, "\n"))
}
