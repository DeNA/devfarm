package devicefarm

import "testing"

func TestNewAuthorizationStatusChecker(t *testing.T) {
	execute := AnySuccessfulExecutor()
	checkAuthorization := NewAuthorizationStatusChecker(execute)

	err := checkAuthorization()

	if err != nil {
		t.Errorf("got %v, want nil", err)
		return
	}
}
