package all

import (
	"github.com/dena/devfarm/internal/pkg/platforms"
	"sync"
)

type FoundDeviceTableEntry struct {
	found         bool
	platformError error
}
type FoundDeviceTable map[platforms.ID]FoundDeviceTableEntry
type DeviceFinderTable map[platforms.ID]platforms.DeviceFinder

func (ps Platforms) FindDevice(device platforms.EitherDevice) FoundDeviceTable {
	result := FoundDeviceTable{}
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for platformID, p := range ps.table {
		wg.Add(1)

		go func(platformID platforms.ID, findDevice platforms.DeviceFinder) {
			defer wg.Done()

			ok, err := findDevice(device)

			mutex.Lock()
			defer mutex.Unlock()
			result[platformID] = FoundDeviceTableEntry{ok, err}
		}(platformID, p.DeviceFinder())
	}

	wg.Wait()
	return result
}
