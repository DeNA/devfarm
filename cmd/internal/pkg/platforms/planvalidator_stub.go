package platforms

import "github.com/dena/devfarm/cmd/internal/pkg/testutil"

func AnyPlanValidator() PlanValidator {
	return StubPlanValidator(testutil.AnyError)
}

func StubPlanValidator(err error) PlanValidator {
	return func(EitherPlan) error {
		return err
	}
}
