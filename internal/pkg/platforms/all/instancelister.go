package all

import (
	"github.com/dena/devfarm/internal/pkg/platforms"
	"sort"
	"sync"
)

type InstancesOrError struct {
	entries       []platforms.InstanceOrError
	platformError error
}
type InstancesListerTable map[platforms.ID]platforms.InstanceLister

func (ps Platforms) ListInstances(groupName platforms.InstanceGroupName) map[platforms.ID]InstancesOrError {
	result := map[platforms.ID]InstancesOrError{}
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for platformID, p := range ps.table {
		wg.Add(1)

		go func(platformID platforms.ID, lister platforms.InstanceLister) {
			defer wg.Done()

			entries, err := lister(groupName)

			mutex.Lock()
			defer mutex.Unlock()
			result[platformID] = InstancesOrError{entries, err}
		}(platformID, p.InstanceLister())
	}

	wg.Wait()
	return result
}

func PlatformInstanceEntryFromTable(table map[platforms.ID]InstancesOrError) []PlatformInstancesListEntry {
	result := make([]PlatformInstancesListEntry, 0)

	for platformID, platformEntry := range table {
		if platformEntry.platformError != nil {
			result = append(result, PlatformInstancesListEntry{
				PlatformID: platformID,
				Entry:      newErrorInstanceEntry(platformEntry.platformError),
			})
			continue
		}

		for _, instanceEntry := range platformEntry.entries {
			result = append(result, PlatformInstancesListEntry{
				PlatformID: platformID,
				Entry:      instanceEntry,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].isLess(result[j])
	})

	return result
}

type PlatformInstancesListEntry struct {
	PlatformID platforms.ID
	Entry      platforms.InstanceOrError
}

func (platformEntry PlatformInstancesListEntry) isLess(another PlatformInstancesListEntry) bool {
	if platformEntry.PlatformID == another.PlatformID {
		return platformEntry.Entry.Less(another.Entry)
	}
	return platformEntry.PlatformID < another.PlatformID
}

func newErrorInstanceEntry(err error) platforms.InstanceOrError {
	return platforms.NewInstanceListEntry(newErrorInstance(), err)
}

func newErrorInstance() platforms.Instance {
	return platforms.NewInstance(
		platforms.NewUnavailableEitherDevice(),
		platforms.InstanceStateIsUnknown,
	)
}
