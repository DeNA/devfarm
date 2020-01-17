package all

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/platforms"
	"github.com/dena/devfarm/cmd/core/platforms/awsdevicefarm"
	"sort"
	"strings"
)

type Platforms struct {
	table map[platforms.ID]platforms.Platform
}

func NewPlatforms(bag platforms.Bag) Platforms {
	return Platforms{table: map[platforms.ID]platforms.Platform{
		awsdevicefarm.ID: awsdevicefarm.NewPlatform(bag),
	}}
}

func (ps Platforms) GetPlatform(id platforms.ID) (platforms.Platform, error) {
	if p, ok := ps.table[id]; ok {
		return p, nil
	}
	return nil, fmt.Errorf("no such platform: %q (available ones are %s)", id, strings.Join(ps.PlatformNames(), ", "))
}

func (ps Platforms) PlatformNames() []string {
	result := make([]string, len(ps.table))

	i := 0
	for platformID := range ps.table {
		result[i] = string(platformID)
		i++
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func (ps Platforms) PlatformIDs() []platforms.ID {
	result := make([]platforms.ID, len(ps.table))

	i := 0
	for platformID := range ps.table {
		result[i] = platformID
		i++
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}
