package all

import (
	"github.com/dena/devfarm/cmd/core/platforms"
	"sort"
	"sync"
)

type DevicesTableEntry struct {
	entries       []platforms.DeviceOrError
	platformError error
}
type DevicesTable map[platforms.ID]DevicesTableEntry
type DevicesListerTable map[platforms.ID]platforms.DeviceLister

func (ps Platforms) ListAllDevices() DevicesTable {
	result := DevicesTable{}
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for platformID, p := range ps.table {
		wg.Add(1)

		go func(platformID platforms.ID, lister platforms.DeviceLister) {
			defer wg.Done()

			entries, err := lister()

			mutex.Lock()
			defer mutex.Unlock()
			result[platformID] = DevicesTableEntry{entries, err}
		}(platformID, p.DeviceLister())
	}

	wg.Wait()
	return result
}

type PlatformDeviceListEntry struct {
	PlatformID platforms.ID
	Entry      platforms.DeviceOrError
}

func (platformEntry PlatformDeviceListEntry) isLess(another PlatformDeviceListEntry) bool {
	if platformEntry.PlatformID == another.PlatformID {
		return platformEntry.Entry.Less(another.Entry)
	}
	return platformEntry.PlatformID < another.PlatformID
}

func DeviceListEntries(table DevicesTable) []PlatformDeviceListEntry {
	var triples []PlatformDeviceListEntry

	for platformID, tableEntry := range table {
		if tableEntry.platformError != nil {
			triples = append(triples, PlatformDeviceListEntry{
				PlatformID: platformID,
				Entry:      platforms.UnspecificErrorDeviceListEntry(tableEntry.platformError),
			})
			continue
		}

		for _, listEntry := range tableEntry.entries {
			triples = append(triples, PlatformDeviceListEntry{
				PlatformID: platformID,
				Entry:      listEntry,
			})
		}
	}

	sort.Slice(triples, func(i, j int) bool {
		return !triples[i].isLess(triples[j])
	})

	return triples
}
