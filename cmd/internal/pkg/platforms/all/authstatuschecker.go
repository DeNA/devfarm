package all

import (
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"sync"
)

func (ps Platforms) CheckAllAuthStatus() map[platforms.ID]error {
	result := map[platforms.ID]error{}
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for platformID, p := range ps.table {
		wg.Add(1)

		go func(platformID platforms.ID, authStatusChecker platforms.AuthStatusChecker) {
			defer wg.Done()

			authStatus := authStatusChecker()

			mutex.Lock()
			defer mutex.Unlock()
			result[platformID] = authStatus
		}(platformID, p.AuthStatusChecker())
	}

	wg.Wait()
	return result
}
