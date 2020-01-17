package platforms

import "testing"

func TestUnoptimizedRunnerEmpty(t *testing.T) {
	runner := NewUnoptimizedRunner(SuccessfulIOSRunner(), SuccessfulAndroidRunner())
	var plans []EitherPlan

	if _, err := runner(plans); err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}
}

func TestUnoptimizedRunnerSuccess(t *testing.T) {
	runner := NewUnoptimizedRunner(SuccessfulIOSRunner(), SuccessfulAndroidRunner())
	plans := []EitherPlan{
		{OSName: OSIsIOS},
		{OSName: OSIsAndroid},
	}

	if _, err := runner(plans); err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}
}

func TestUnoptimizedRunnerIOSFailure(t *testing.T) {
	runner := NewUnoptimizedRunner(FailedIOSRunner(), SuccessfulAndroidRunner())
	plans := []EitherPlan{
		{OSName: OSIsIOS},
		{OSName: OSIsAndroid},
	}

	if _, err := runner(plans); err == nil {
		t.Error("want error, got nil")
		return
	}
}

func TestUnoptimizedRunnerAndroidFailure(t *testing.T) {
	runner := NewUnoptimizedRunner(SuccessfulIOSRunner(), FailedAndroidRunner())
	plans := []EitherPlan{
		{OSName: OSIsIOS},
		{OSName: OSIsAndroid},
	}

	if _, err := runner(plans); err == nil {
		t.Error("want error, got nil")
		return
	}
}

func TestUnoptimizedRunnerBothFailure(t *testing.T) {
	runner := NewUnoptimizedRunner(FailedIOSRunner(), FailedAndroidRunner())
	plans := []EitherPlan{
		{OSName: OSIsIOS},
		{OSName: OSIsAndroid},
	}

	if _, err := runner(plans); err == nil {
		t.Error("want error, got nil")
		return
	}
}
