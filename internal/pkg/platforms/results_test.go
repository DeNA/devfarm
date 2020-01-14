package platforms

import (
	"github.com/dena/devfarm/internal/pkg/testutil"
	"testing"
)

func TestResultsSuccess(t *testing.T) {
	_, err := successfulResultsGen()

	if err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}
}

func TestResultsFailed(t *testing.T) {
	_, err := failedResultsGen()

	if err == nil {
		t.Error("want error, got nil")
		return
	}
}

func successfulResultsGen() (Results, error) {
	results := NewResults(nil)
	return *results, results.Err()
}

func failedResultsGen() (Results, error) {
	results := NewResults(testutil.AnyError)
	return *results, results.Err()
}
