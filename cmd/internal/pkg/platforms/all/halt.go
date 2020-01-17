package all

import (
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"sync"
)

func (ps Platforms) HaltAll(groupName platforms.InstanceGroupName) (ResultTable, error) {
	var wg sync.WaitGroup
	builder := NewResultTableBuilder()
	for platformID, p := range ps.table {
		wg.Add(1)

		go func(platformID platforms.ID, halt platforms.Halt) {
			defer wg.Done()
			results, _ := halt(groupName)
			builder.AddErrors(platformID, results.ErrorsIncludingNotNil()...)
		}(platformID, p.Halt())
	}

	wg.Wait()
	table := builder.Build()
	return table, table.Err()
}
