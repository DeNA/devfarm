package all

import (
	"github.com/dena/devfarm/internal/pkg/platforms"
	"sync"
)

func HaltAll(groupName platforms.InstanceGroupName, bag platforms.HaltBag) (ResultTable, error) {
	platformTable := make(map[platforms.ID]platforms.Halt, len(PlatformTable))

	for _, platform := range PlatformTable {
		platformTable[platform.ID()] = platform.Halt()
	}

	return haltAllOn(platformTable, groupName, bag)
}

// Injectable version for testing.
func haltAllOn(platformTable map[platforms.ID]platforms.Halt, groupName platforms.InstanceGroupName, bag platforms.HaltBag) (ResultTable, error) {
	var wg sync.WaitGroup
	builder := NewResultTableBuilder()
	for platformID, halt := range platformTable {
		wg.Add(1)

		go func(platformID platforms.ID, halt platforms.Halt) {
			defer wg.Done()
			results, _ := halt(groupName, bag)
			builder.AddError(platformID, results.ErrorsIncludingNotNil()...)
		}(platformID, halt)
	}

	wg.Wait()
	table := builder.Build()
	return table, table.Err()
}
