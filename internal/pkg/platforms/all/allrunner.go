package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"sync"
)

func RunAll(plans []platforms.EitherPlan, bag platforms.Bag) (ResultTable, error) {
	if errTable, err := ValidatePlans(bag, plans); err != nil {
		return errTable, err
	}

	builder := NewResultTableBuilder()
	var wg sync.WaitGroup

	for _, plan := range plans {
		platform, platformErr := GetPlatform(plan)
		if platformErr != nil {
			builder.AddError(plan.Platform, platformErr)
			continue
		}

		wg.Add(1)
		go func(platform platforms.Platform, plan platforms.EitherPlan) {
			var err error

			switch plan.OSName {
			case platforms.OSIsIOS:
				runIOS := platform.IOSRunner()
				err = runIOS(plan.IOS(), bag)

			case platforms.OSIsAndroid:
				runAndroid := platform.AndroidRunner()
				err = runAndroid(plan.Android(), bag)

			default:
				err = fmt.Errorf("not supported OS: %q", plan.OSName)
			}

			builder.AddError(platform.ID(), err)
			wg.Done()
		}(platform, plan)
	}
	wg.Wait()

	table := builder.Build()
	return table, table.Err()
}
