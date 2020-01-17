package all

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/platforms"
	"sync"
)

func (ps Platforms) ValidatePlans(plans []platforms.EitherPlan) (ResultTable, error) {
	var wg sync.WaitGroup
	builder := NewResultTableBuilder()

	for _, plan := range plans {
		platform, platformErr := ps.GetPlatform(plan.Platform)
		if platformErr != nil {
			builder.AddErrors(plan.Platform, fmt.Errorf("no such platform: %q", plan.Platform))
			continue
		}

		wg.Add(1)
		go func(plan platforms.EitherPlan, platform platforms.Platform) {
			validatePlan := platform.PlanValidator()
			if err := validatePlan(plan); err != nil {
				builder.AddErrors(platform.ID(), err)
			} else {
				builder.AddErrors(platform.ID(), nil)
			}
			wg.Done()
		}(plan, platform)
	}

	wg.Wait()
	table := builder.Build()
	return table, table.Err()
}
