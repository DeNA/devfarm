package platforms

import (
	"fmt"
	"strings"
)

// Result is a struct represents that operations for an single instance was succeeded or not.
//
// NOTE: Devfarm may control multiple instances, and controlling some of them may be failed.
// At the point, there are some way to represent how many instances successfully controlled,
// and how many errors occurred and why the errors occurred.
//
// In this project, we represents them by ([]Result, error), and typically the error is a *SomeErrors.
// A Result in []Result knows whether operations for a single instance were succeeded,
// or a failure reason if the operations were failed.
//
// ([]Result, error) has an advantage that developers can know whether all operations were succeeded or not by
// Go's standard convention of error detection. And can know how many operations are failed and why some of
// operations are failed by reading Result.Err().
type Results []error

func NewResults(errs ...error) *Results {
	r := make(Results, len(errs))
	for i, err := range errs {
		r[i] = err
	}
	return &r
}

func (r *Results) AddErrorOrNils(errs ...error) {
	*r = append(*r, errs...)
}

func (r *Results) AddSuccesses(delta int) {
	var i int
	for i = 0; i < delta; i++ {
		*r = append(*r, nil)
	}
}

func (r *Results) ErrorsIncludingNotNil() []error {
	return *r
}

func (r *Results) ErrorsOnlyNotNil() []error {
	errs := make([]error, 0)
	for _, err := range *r {
		if err == nil {
			continue
		}
		errs = append(errs, err)
	}
	return errs
}

func (r *Results) Err() error {
	errs := r.ErrorsOnlyNotNil()
	count := len(errs)
	if count > 0 {
		messages := make([]string, count)
		for i, err := range errs {
			messages[i] = err.Error()
		}
		return fmt.Errorf("%d errors ocurred:\n%s", count, strings.Join(messages, "\n"))
	}
	return nil
}
