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

func FindDevice(device platforms.EitherDevice, bag platforms.DeviceFinderBag) FoundDeviceTable {
	finderTable := make(DeviceFinderTable, len(PlatformTable))

	for _, platform := range PlatformTable {
		finderTable[platform.ID()] = platform.DeviceFinder()
	}

	return findDeviceOn(finderTable, device, bag)
}

func findDeviceOn(finderTable DeviceFinderTable, device platforms.EitherDevice, bag platforms.DeviceFinderBag) FoundDeviceTable {
	result := FoundDeviceTable{}
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for platformID, finder := range finderTable {
		wg.Add(1)

		go func(platformID platforms.ID, findDevice platforms.DeviceFinder) {
			defer wg.Done()

			ok, err := findDevice(device, bag)

			mutex.Lock()
			defer mutex.Unlock()
			result[platformID] = FoundDeviceTableEntry{ok, err}
		}(platformID, finder)
	}

	wg.Wait()
	return result
}
