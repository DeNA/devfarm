package all

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/platforms"
	"github.com/dena/devfarm/cmd/core/platforms/awsdevicefarm"
	"strings"
)

var ValidPlatformIDs = map[platforms.ID]bool{
	awsdevicefarm.ID: true,
}

func ValidatePlatformID(platformName string) (platforms.ID, error) {
	unsafeId := platforms.ID(platformName)

	if _, ok := ValidPlatformIDs[unsafeId]; ok {
		safeId := unsafeId
		return safeId, nil
	}

	availableOnes := make([]string, len(ValidPlatformIDs))
	i := 0
	for available := range ValidPlatformIDs {
		availableOnes[i] = string(available)
		i++
	}
	return "", fmt.Errorf("no such platform: %q (available ones are %s)", platformName, strings.Join(availableOnes, ", "))
}
