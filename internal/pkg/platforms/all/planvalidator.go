package all

import (
	"github.com/dena/devfarm/internal/pkg/platforms"
	"sync"
)

func ValidatePlans(bag platforms.PlanValidatorBag, plans []platforms.EitherPlan) (ResultTable, error) {
	var wg sync.WaitGroup
	builder := NewResultTableBuilder()

	for _, plan := range plans {
		platform, platformErr := GetPlatform(plan)
		if platformErr != nil {
			continue
		}

		wg.Add(1)
		go func(plan platforms.EitherPlan, platform platforms.Platform) {
			validatePlan := platform.PlanValidator()
			err := validatePlan(bag, plan)
			builder.AddError(platform.ID(), err)
			wg.Done()
		}(plan, platform)
	}

	wg.Wait()
	table := builder.Build()
	return table, table.Err()
}
