package platforms

import "testing"

func TestUnoptimizedForeverEmpty(t *testing.T) {
	runner := NewUnoptimizedForever(SuccessfulIOSForever(), SuccessfulAndroidForever())
	var plans []EitherPlan

	_, err := runner(plans)

	if err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}
}

func TestUnoptimizedForeverSuccess(t *testing.T) {
	runner := NewUnoptimizedForever(SuccessfulIOSForever(), SuccessfulAndroidForever())
	plans := []EitherPlan{
		{OSName: OSIsIOS},
		{OSName: OSIsAndroid},
	}

	_, err := runner(plans)

	if err != nil {
		t.Errorf("want nil, got %v", err)
		return
	}
}

func TestUnoptimizedForeverIOSFailure(t *testing.T) {
	runner := NewUnoptimizedForever(FailedIOSForever(), SuccessfulAndroidForever())
	plans := []EitherPlan{
		{OSName: OSIsIOS},
		{OSName: OSIsAndroid},
	}

	_, err := runner(plans)

	if err == nil {
		t.Error("want error, got nil")
		return
	}
}

func TestUnoptimizedForeverAndroidFailure(t *testing.T) {
	runner := NewUnoptimizedForever(SuccessfulIOSForever(), FailedAndroidForever())
	plans := []EitherPlan{
		{OSName: OSIsIOS},
		{OSName: OSIsAndroid},
	}

	_, err := runner(plans)

	if err == nil {
		t.Error("want error, got nil")
		return
	}
}

func TestUnoptimizedForeverBothFailure(t *testing.T) {
	runner := NewUnoptimizedForever(FailedIOSForever(), FailedAndroidForever())
	plans := []EitherPlan{
		{OSName: OSIsIOS},
		{OSName: OSIsAndroid},
	}

	_, err := runner(plans)

	if err == nil {
		t.Error("want error, got nil")
		return
	}
}
