package all

import (
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"sync"
)

func (ps Platforms) ListAllInstances() map[platforms.ID]InstancesOrError {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	table := make(map[platforms.ID]InstancesOrError, len(ps.table))

	for platformID, p := range ps.table {
		wg.Add(1)
		go func(platformID platforms.ID, listAllInstances platforms.AllInstanceLister) {
			defer wg.Done()

			entries, err := listAllInstances()

			mutex.Lock()
			defer mutex.Unlock()
			table[platformID] = InstancesOrError{entries, err}
		}(platformID, p.AllInstanceLister())
	}

	wg.Wait()
	return table
}
