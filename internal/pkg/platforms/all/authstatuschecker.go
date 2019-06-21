package all

import (
	"github.com/dena/devfarm/internal/pkg/platforms"
	"sync"
)

func CheckAllAuthStatus(bag platforms.AuthStatusCheckerBag) map[platforms.ID]error {
	platformCheckerTable := make(map[platforms.ID]platforms.AuthStatusChecker, len(PlatformTable))

	for _, platform := range PlatformTable {
		platformCheckerTable[platform.ID()] = platform.AuthStatusChecker()
	}

	return CheckAuthStatusOn(platformCheckerTable, bag)
}

func CheckAuthStatusOn(table map[platforms.ID]platforms.AuthStatusChecker, bag platforms.AuthStatusCheckerBag) map[platforms.ID]error {
	result := map[platforms.ID]error{}
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for platformID, authStatusChecker := range table {
		wg.Add(1)

		go func(platformID platforms.ID, authStatusChecker platforms.AuthStatusChecker) {
			defer wg.Done()

			authStatus := authStatusChecker(bag)

			mutex.Lock()
			defer mutex.Unlock()
			result[platformID] = authStatus
		}(platformID, authStatusChecker)
	}

	wg.Wait()
	return result
}
