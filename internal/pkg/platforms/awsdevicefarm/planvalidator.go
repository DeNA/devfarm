package awsdevicefarm

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"time"
)

// There is a 150-minute limit to the duration of an automated test run.
// https://docs.aws.amazon.com/en_pv/devicefarm/latest/developerguide/limits.html
const maxLifetime = 150 * time.Minute

func newPlanValidator() platforms.PlanValidator {
	return func(bag platforms.PlanValidatorBag, plan platforms.EitherPlan) error {
		lifetime := plan.CommonPart.Lifetime
		if lifetime > maxLifetime {
			return fmt.Errorf("liftime must be shorter than 150 minutes: got %.0f minutes", lifetime.Minutes())
		}
		return nil
	}
}
