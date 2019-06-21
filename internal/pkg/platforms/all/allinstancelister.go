package all

import (
	"github.com/dena/devfarm/internal/pkg/platforms"
	"sync"
)

func ListAllInstances(bag platforms.AllInstanceListerBag) map[platforms.ID]InstancesOrError {
	platformTable := make(map[platforms.ID]platforms.AllInstanceLister, len(PlatformTable))

	for _, platform := range PlatformTable {
		platformTable[platform.ID()] = platform.AllInstanceLister()
	}

	return listAllInstancesOn(platformTable, bag)
}

func listAllInstancesOn(platformTable map[platforms.ID]platforms.AllInstanceLister, bag platforms.AllInstanceListerBag) map[platforms.ID]InstancesOrError {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	table := make(map[platforms.ID]InstancesOrError, len(PlatformTable))

	for platformID, listAllInstances := range platformTable {
		wg.Add(1)
		go func(platformID platforms.ID, listAllInstances platforms.AllInstanceLister) {
			defer wg.Done()

			entries, err := listAllInstances(bag)

			mutex.Lock()
			defer mutex.Unlock()
			table[platformID] = InstancesOrError{entries, err}
		}(platformID, listAllInstances)
	}

	wg.Wait()

	return table
}
