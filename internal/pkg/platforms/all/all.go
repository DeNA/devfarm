package all

import (
	"fmt"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"github.com/dena/devfarm/internal/pkg/platforms/awsdevicefarm"
	"sort"
)

var PlatformTable = map[platforms.ID]platforms.Platform{
	awsdevicefarm.ID: awsdevicefarm.AWSDeviceFarm,
}

func GetPlatform(plan platforms.EitherPlan) (platforms.Platform, error) {
	if platform, ok := PlatformTable[plan.CommonPart.Platform]; ok {
		return platform, nil
	}
	return nil, fmt.Errorf("no such platform: %q", plan.CommonPart.Platform)
}

func PlatformNames() []string {
	result := make([]string, len(PlatformTable))

	i := 0
	for platformID := range PlatformTable {
		result[i] = string(platformID)
		i++
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func PlatformIDs() []platforms.ID {
	result := make([]platforms.ID, len(PlatformTable))

	i := 0
	for platformID := range PlatformTable {
		result[i] = platformID
		i++
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}
