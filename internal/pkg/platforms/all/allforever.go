package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"sync"
)

func ForeverAll(plans []platforms.EitherPlan, bag platforms.Bag) (ResultTable, error) {
	builder := NewResultTableBuilder()
	var wg sync.WaitGroup

	for _, plan := range plans {
		wg.Add(1)
		go func(plan platforms.EitherPlan) {
			defer wg.Done()
			var err error

			switch plan.OSName {
			case platforms.OSIsIOS:
				err = ForeverIOS(plan.IOS(), bag)
			case platforms.OSIsAndroid:
				err = ForeverAndroid(plan.Android(), bag)
			default:
				err = fmt.Errorf("not supported OS: %q", plan.OSName)
			}

			builder.AddError(plan.Platform, err)
		}(plan)
	}

	wg.Wait()
	table := builder.Build()
	return table, table.Err()
}
